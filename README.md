# Pentimento
> In painting, a pentimento (italian) is "the presence or emergence of earlier images, forms, or strokes that have been changed and painted over"


This project aims to explore, research, and implement some techniques in the field of image steganography using Golang.

## What is steganography?
![stenography](assets/SteganographicModel.png)

## LSB - basic technique used for steanography
>
The Least Significant Bit (LSB) is a method used in digital steganography for hiding information within a digital file, such as an image or audio file.

In the context of an image, each pixel is represented by a binary number. The "least significant bit" is the last bit in this binary representation. The LSB method works by replacing these least significant bits with the bits from the data that needs to be hidden.

Because the least significant bit has the smallest impact on the overall value, changing it usually results in a minor alteration to the pixel's color. This change is typically so small that it's imperceptible to the human eye, making LSB a popular method for hiding information within images.

![LSB](assets/LSB.jpg)


## Useful papers 


### [Highly Secured Hybrid Image Steganography with an Improved Key](https://dergipark.org.tr/tr/download/article-file/2475349)
Generation and Exchange for One-Time-Pad Encryption Method



Summary from the paper
```
Discrete Haar Wavelet Transform (DHWT)

One-Time-Pad (OTP) Encryption

Highly Secured Information Exchange Algorithm (HSIEA)

Least Significant Bit (LSB) Method

Optimal Pixel Adjustment Process (OPAP)

Discrete Cosine Transform (DCT)
```


### [IMAGE BASED STEGANOGRAPHY AND CRYPTOGRAPHY](https://www.diag.uniroma1.it/~bloisi/steganography/isc.pdf)



They propose a new method for integrating cryptography and steganography, which they call ISC (Image-based Steganography and Cryptography). 

It uses images as cover objects for steganography and as keys for cryptography. 

It's designed to work with bit streams scattered over multiple images or with still images. The method yields random outputs to make steganalysis more difficult and can cipher the message in a theoretically secure manner while preserving the stego image's statistical properties.




### [The Art of Data Hiding with Reed-Solomon Error Correcting Codes](https://arxiv.org/abs/1411.4790)



Basically this paper discusses the use of Reed-Solomon error correcting codes in steganography, which is the art of hiding information in a way that is not detectable to the naked eye. The authors propose a design that substitutes redundant Reed-Solomon codes with the steganographic message. 

### Steganography Toolkits
```
docker run -it --rm -v $(pwd)/data:/data dominicbreuker/stego-toolkit /bin/bash
zsteg -a out.png
```
For now LSB is easy to detect

### Name idea
https://en.wikipedia.org/wiki/Pentimento


