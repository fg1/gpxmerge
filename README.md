gpxmerge
========

`gpxmerge` is a small tool to merge multiple GPX tracks in a single one.

The program is currently in an alpha state, without particular effort on input GPX validation/sanitization/etc. or customizability.


## Compilation and installation

On Debian/Ubuntu:
```
# apt-get install golang
$ git clone https://github.com/fg1/gpxmerge.git
$ cd gpxmerge
$ go build
```


## Usage
```
gpxmerge -o out.gpx inp1.gpx inp2.gpx ...
```


## Contributing

Contributions are welcome.

1. [Fork the repository](https://github.com/fg1/gpxmerge/fork)
2. Create your feature branch (`git checkout -b my-feature`)
3. Commit your changes (`git commit -am 'Commit message'`)
4. Push to the branch (`git push origin my-feature`)
5. Create a pull request

