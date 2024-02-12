# WordListCombinator
[![Go](https://github.com/PatchRequest/WordListCombinator/actions/workflows/go.yml/badge.svg)](https://github.com/PatchRequest/WordListCombinator/actions/workflows/go.yml)  
Patchy's WordListCombinator combines multiples Wordlists fast and with an ultra low memory footprint. The combined wordlist will not contain any duplicates. A direct append of data to the wordlist in cracking tools such as Hashcat, or similar tools, is inefficient and wasteful of resources because these tools do not deduplicate the wordlist.
The efficiency is  achieved by using bloomfilters, a space-efficient probabilistic data structure.

# Installation
WordListCombinator is an all-in-one binary just download it on the release page [here](https://link-url-here.org)

# How to use
```
./WordListCombinator -receiver all_in_one_w.txt -sender hashkiller-dict.txt -receiversize 1184801557 -fprate 0.01
```
- Receiver:   Wordlist which will get appended with new entries (Default: receiver.txt)
- Sender: Wordlist which will be included in the receiver after run (Default: sender.txt)
- Receiversize: Line count of the receiver wordlist (Default: 10,000,000)
- fprate: Probability of false positive and therefore false rejection. lower fprate -> higher memory footprint (Default: 0.01 = 1%)  

# Benchmarks
My very unscientific benchmark on my M2 Mac Mini
```
[*] File Receiver: 
All-in-One-Wi-Fi
134.17GB
11848015579 Lines

[*] File Sender:
Hashkiller-dict
2.76GB
266517972 Lines

[*] Result:
9749.22s user 1225.97s system 383% cpu 47:39.85 total
<5 GB RAM used
+307424 new Unique Lines
11848323003 Lines Total
```
If you ran it for yourself on different hardware feel free to open a pull request to add your benchmark

# Contributing
Pull requests are welcome. Feel free to open an issue if you want to add other features.

# Credits
- Bloomfilter Library: https://github.com/bits-and-blooms/bloom/
- Burton Howard Bloom for the creation of bloomfilters
