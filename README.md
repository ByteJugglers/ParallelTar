# ParallelTar 并行打包

This is a simple small program, mainly used to compress and decompress files in a specified directory in batches according to the specified compression format.
这是一个简易的小程序，主要用于将指定目录下的文件按照指定的压缩格式批量的压缩和解压。

## Example 示例
```shell
$ ./parallel_tar --help
NAME:
   parallel_tar - Tar files parallel.

USAGE:
   ParallelTar [global options] command [command options]

VERSION:
   1.0.0

COMMANDS:
   tar      tar files or directories
   untar    Unzip files in the directory
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

# Batch Compression
$ ./parallel_tar tar -d data -t gzip -j 4
2024/06/22 19:18:35 Packing la32r-Linux.tar.gz completed.
2024/06/22 19:18:46 Packing linux_old1.tar.gz completed.

# Batch decompression
$ ./parallel_tar untar -d data -j 4
2024/06/22 19:18:10 Unpacking kicad-packages3D.tar.gz completed.
2024/06/22 19:18:10 Unpacking koa-mongo-ts-api.tar.gz completed.
2024/06/22 19:18:13 Unpacking kernel.tar.gz completed.
```