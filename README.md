# Go ProtoBuf

This repository is the same representation of the [Java Protobuf example](https://github.com/davidlares/java-protobuf) but with `GoLang`. The `src/simple` directory contains the `simple.proto` and the outputted version resulting from the `protoc` compilation process.

The `main.go` file is a simple implementation of a `Protobuf Message`, that scripts shows you how can be translated and manipulated with `Go`. The `import` section contains the script dependencies.

The `protoc` commands used are in the `generate.bat` (Windows-based) and `generate.sh` (Linux-based).

## Let's recap a bit

The Protocol buffers are a data-type (actually a format) developed by Google used for sharing data across many development languages. This particular data-type is structured and it's defined with the `.proto` file extension, also, is fully typed, very comprehensive but schema-based.

Some databases offer support for `Protobuf` data format, this tool is internally used by Google and a lot of RPC frameworks have a natural adoption in order to exchange data. You can check for more information [here](https://developers.google.com/protocol-buffers/).

## Usage

`go run main.go` (you can also `build` the script and then run it as well)

## Credits

 [David E Lares](https://twitter.com/davidlares3)

## License

 [MIT](https://opensource.org/licenses/MIT)
