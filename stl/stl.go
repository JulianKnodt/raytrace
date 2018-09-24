package stl

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"sync"

	v "github.com/julianknodt/vector"
)

var mantessaFmt = `(-?\d+.?\d+(e[-+]?(\d+))?)`
var Header = regexp.MustCompile(`solid\s+(.+)`)
var FacetNormal = regexp.MustCompile(fmt.Sprintf(`facet normal\s+%s\s+%s\s+%s`,
	mantessaFmt, mantessaFmt, mantessaFmt))

var EndFacet = regexp.MustCompile("endfacet")
var EndLoop = regexp.MustCompile("endloop")
var OuterLoop = regexp.MustCompile(`outer\s+loop`)
var Vertex = regexp.MustCompile(fmt.Sprintf(`vertex\s+%s\s+%s\s+%s`,
	mantessaFmt, mantessaFmt, mantessaFmt))

var EndSolid = regexp.MustCompile("endsolid")

// Minimal Implementation of Triangle for STLs

type Triangle struct {
	// The normal to the triangle
	Normal v.Vec3

	// The 3 points of the triangle
	V1, V2, V3 v.Vec3
}

func MatchesNormalOrEnd(rd *bufio.Reader) (vec v.Vec3, err error) {
	var normalInfo string
	normalInfo, err = rd.ReadString('\n')
	if FacetNormal.MatchString(normalInfo) {
		matches := FacetNormal.FindAllStringSubmatch(normalInfo, -1)
		n1, err1 := strconv.ParseFloat(matches[0][1], 64)
		n2, err2 := strconv.ParseFloat(matches[0][4], 64)
		n3, err3 := strconv.ParseFloat(matches[0][7], 64)
		switch {
		case err1 != nil:
			err = err1
		case err2 != nil:
			err = err2
		case err3 != nil:
			err = err3
		}
		vec = v.Vec3{n1, n2, n3}
	} else if EndSolid.MatchString(normalInfo) {
		// Handle here because we can't determine when the end is
		err = io.EOF
	} else if err == nil {
		err = fmt.Errorf("Cannot parse %s", normalInfo)
	}
	return
}

func MatchesVertex(rd *bufio.Reader) (vec v.Vec3, err error) {
	var vertexInfo string
	vertexInfo, err = rd.ReadString('\n')
	if Vertex.MatchString(vertexInfo) {
		matches := Vertex.FindAllStringSubmatch(vertexInfo, -1)
		// 1, 4, 7 because of subgroups matching
		v1, err1 := strconv.ParseFloat(matches[0][1], 64)
		v2, err2 := strconv.ParseFloat(matches[0][4], 64)
		v3, err3 := strconv.ParseFloat(matches[0][7], 64)
		switch {
		case err1 != nil:
			err = err1
		case err2 != nil:
			err = err2
		case err3 != nil:
			err = err3
		}
		vec = v.Vec3{v1, v2, v3}
	} else if err == nil {
		err = fmt.Errorf("Cannot parse %s", vertexInfo)
	}
	return
}

func MatchesEndLoop(rd *bufio.Reader) error {
	el, err := rd.ReadString('\n')
	switch {
	case err != nil:
	case !EndLoop.MatchString(el):
		err = fmt.Errorf("Invalid end loop string %s", el)
	}
	return err
}

func MatchesEndFacet(rd *bufio.Reader) error {
	el, err := rd.ReadString('\n')
	switch {
	case err != nil:
	case !EndFacet.MatchString(el):
		err = fmt.Errorf("Invalid end facet string %s", el)
	}
	return err
}

func DecodeAscii(r io.Reader) (o []Triangle, err error) {
	rd := bufio.NewReader(r)
	var header string
	header, err = rd.ReadString('\n')
	switch {
	case err != nil:
		return
	case !Header.MatchString(header):
		err = fmt.Errorf("Invalid header format %s", header)
		return
	}

	o = make([]Triangle, 0, 64)
	for err == nil {
		t := Triangle{}
		var n v.Vec3
		n, err = MatchesNormalOrEnd(rd)
		if err != nil {
			break
		}
		t.Normal = n

		var ol string
		ol, err = rd.ReadString('\n')
		if err != nil {
			break
		} else if !OuterLoop.MatchString(ol) {
			err = fmt.Errorf("outer loop was mismatched with %s", ol)
			break
		}

		v1, err := MatchesVertex(rd)
		t.V1 = v1
		v2, err := MatchesVertex(rd)
		t.V2 = v2
		v3, err := MatchesVertex(rd)
		t.V3 = v3

		o = append(o, t)
		if err = MatchesEndLoop(rd); err != nil {
			break
		}

		if err = MatchesEndFacet(rd); err != nil {
			break
		}
	}

	if err == io.EOF {
		err = nil
	}
	return
}

const BinaryHeaderSize = 80

var infoOnce sync.Once

func DecodeBinary(r io.Reader) (o []Triangle, err error) {
	// Ignore first 80 bytes
	if _, err = io.CopyN(ioutil.Discard, r, BinaryHeaderSize); err != nil {
		return
	}

	buf := make([]byte, 4)
	if _, err = r.Read(buf); err != nil {
		return
	}
	bleu := binary.LittleEndian.Uint32
	fl32 := math.Float32frombits
	numTriangles := binary.LittleEndian.Uint32(buf)
	o = make([]Triangle, 0, numTriangles)

	vb := make([]byte, 50)
	for i := uint32(0); i < numTriangles; i++ {
		_, err = r.Read(vb)
		switch {
		case err == io.EOF:
		case err != nil:
			return
		}
		t := Triangle{}
		t.V1 = *v.VecFloat32(fl32(bleu(vb[12:16])), fl32(bleu(vb[16:20])), fl32(bleu(vb[20:24])))
		t.V2 = *v.VecFloat32(fl32(bleu(vb[24:28])), fl32(bleu(vb[28:32])), fl32(bleu(vb[32:36])))
		t.V3 = *v.VecFloat32(fl32(bleu(vb[36:40])), fl32(bleu(vb[40:44])), fl32(bleu(vb[44:48])))
		t.Normal = *v.VecFloat32(fl32(bleu(vb[:4])), fl32(bleu(vb[4:8])), fl32(bleu(vb[8:12])))
		if t.Normal.IsZero() {
			t.Normal = *t.V1.Sub(t.V2).CrossSet(*t.V1.Sub(t.V3)).UnitSet()
		}
		attributeByteCount := binary.LittleEndian.Uint16(vb[48:50])
		if attributeByteCount != 0 {
			infoOnce.Do(func() {
				fmt.Printf("Cannot determine how to handle %d attribute byte count, ignoring...\n",
					attributeByteCount)
			})
		}
		o = append(o, t)
	}

	return
}
