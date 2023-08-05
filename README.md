# Pentimento
> In painting, a pentimento (italian) is "the presence or emergence of earlier images, forms, or strokes that have been changed and painted over"


This project aims to explore, research, and implement some techniques in the field of image steganography using Golang.

## What is steganography?
![stenography](assets/SteganographicModel.png)

## LSB - basic technique used for steanography
![LSB](assets/LSB.jpg)

## Few useful papers about stenography


### Highly Secured Hybrid Image Steganography with an Improved Key
Generation and Exchange for One-Time-Pad Encryption Method

https://dergipark.org.tr/tr/download/article-file/2475349

summary from the paper
Discrete Haar Wavelet Transform (DHWT)
One-Time-Pad (OTP) Encryption
Highly Secured Information Exchange Algorithm (HSIEA)
Least Significant Bit (LSB) Method
Optimal Pixel Adjustment Process (OPAP)
Discrete Cosine Transform (DCT)


### IMAGE BASED STEGANOGRAPHY AND CRYPTOGRAPHY

https://www.diag.uniroma1.it/~bloisi/steganography/isc.pdf

They propose a new method for integrating cryptography and steganography, which they call ISC (Image-based Steganography and Cryptography). 
It uses images as cover objects for steganography and as keys for cryptography. 
It's designed to work with bit streams scattered over multiple images or with still images. The method yields random outputs to make steganalysis more difficult and can cipher the message in a theoretically secure manner while preserving the stego image's statistical properties.


### The Art of Data Hiding with Reed-Solomon Error Correcting Codes

https://arxiv.org/abs/1411.4790
Basically this paper discusses the use of Reed-Solomon error correcting codes in steganography, which is the art of hiding information in a way that is not detectable to the naked eye. The authors propose a design that substitutes redundant Reed-Solomon codes with the steganographic message. 

### Name idea
https://en.wikipedia.org/wiki/Pentimento


