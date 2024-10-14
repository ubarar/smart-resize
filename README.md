# smart-resize

Use `libvips` to resize images for [pig.js](https://github.com/schlosser/pig.js/).

- Only works on `jpeg`
- Preserves 99% of the original quality
- Does not use chroma subsampling

### Why?

- I wanted something intelligently determine which images need to be resized
- My first attempt to do so in python didn't seem fast enough
- Golang's default `image` library always encodes in 4:2:0

### Results

On my benchmark of 40 files, this program runs about 30% faster than the same thing written in python