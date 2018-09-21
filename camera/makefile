test: clean
	go test -bench=. -benchmem -cpuprofile=cam_prof.out

clean:
	-rm camera.test cam_prof.out
	-rm *.svg *.png
