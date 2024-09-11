# GoCrypt - Cryptographic Framework in Go

GoCrypt is a lightweight cryptographic framework written in Go, designed for developers and security enthusiasts. It offers a variety of cryptographic utilities such as hashing, password cracking, ROT conversion, and base conversion. Whether you're exploring cryptography or building secure applications, GoCrypt provides simple and efficient tools to handle these tasks.

# Features

1) Hashing: Generate cryptographic hash functions including MD5, SHA-1, SHA-256, and others.
2) Password Cracker: A basic password-cracking tool that supports dictionary attacks to test hash resilience.
3) ROT Conversion: Perform ROT13 and ROT-N conversions for encoding/decoding.
4) Base Conversion: Convert between different base systems (binary, decimal, hexadecimal, etc.).

# Getting Started

1) Clone the repository:
 <pre>git clone https://github.com/Fakechippies/Gocrypt.git</pre>

# Usage
<pre>go run main.go [options] [input]</pre>

# Examples
<pre>go run main.go -crack md5 -w /usr/share/wordlists/rockyou.txt 5f4dcc3b5aa765d61d8327deb882cf99</pre>
<pre>go run main.go -hash sha512 GoCrypt</pre>
<pre>go run main.go -enc base64 -e Gocrypt</pre>
<pre>go run main.go -base -f base10 -t base2 777</pre>
<pre>go run main.go -rot 13 -e hello</pre>
 
# Contributing
Contributions are welcome! Feel free to submit pull requests or raise issues.
