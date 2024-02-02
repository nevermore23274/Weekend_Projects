##Build
```
docker build -t meep:latest -f docker/Dockerfile .
```

##Run
```
docker run --name meep -v <pwd>:/home/workdir -it meep
```
- Or if you've ran the container before:
```
docker run -v <pwd>:/home/workdir -it meep
```