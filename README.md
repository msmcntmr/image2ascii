# Image to ASCII Converter
This Go program converts an input image into an ASCII art representation. It uses a character set to replace the pixels in the image with characters based on brightness, generating an ASCII-styled version of the original image.

## Features
- Supports PNG, JPEG, and GIF image formats.
- Converts the image into ASCII art using a customizable character set.
- Automatically resizes the image based on the character dimensions.
- Outputs the resulting ASCII art image in the same format as the input image.

## Prerequisites
- Go 1.18 or higher.
- A valid .ttf font file (for customizing the font).

## Installation
1. Clone this repository:

```bash
git clone https://github.com/msmcntmr/image2ascii.git
cd image2ascii
```
2. Install dependencies:

```bash
go mod tidy
```
Ensure that you have a compatible Go environment set up.

## Usage
1. Prepare Your Image:
    - Ensure your image is in PNG, JPEG, or GIF format.
2. Run the Program:
    ```bash
    go run main.go <filename>
    ```
    Replace `<filename>` with the path to your image file.
    Example:

    ```bash
    go run main.go example.png
    ```
3. Output:
    - The processed image will be saved as `<filename>_processed.<ext>` in the same directory as the input image.

## Example
```bash
go run main.go input.jpg
```
- This will generate a file named input_processed.jpg with the ASCII art version of the image.


## Customization
You can modify the following aspects of the program:

- __Character Set__: The charset variable contains the characters used for ASCII conversion. Modify this string to customize the characters used.
- __Font__: The program uses the `basicfont.Face7x13` font for rendering text. You can replace this with a custom font by loading a .ttf font file.
- __Character Dimensions__: The `charW` and `charH` variables define the width and height of each character block. Adjust these values to fit your needs.

## How it Works
1. __Image Resizing__: The image is resized to match the dimensions of the ASCII grid, determined by the character width and height.
2. __Brightness Calculation__: For each pixel in the resized image, the brightness is calculated by averaging the red, green, and blue components.
3. __Character Mapping__: The brightness value is used to map the pixel to a corresponding character from the charset.
4. __ASCII Art Rendering__: The characters are rendered onto a new image, preserving the original colors but replacing pixels with ASCII characters.

## Limitations
- Small images may result in overly compressed or unclear ASCII art. You can adjust the font size or the character dimensions to improve the output.
- The program currently supports a fixed set of font dimensions. You can modify the code to use a dynamic font size based on the image dimensions.