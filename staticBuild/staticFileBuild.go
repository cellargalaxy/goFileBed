package staticBuild

var Base64Map = map[string]string{}

func init() {
	Base64Map["goFileBed.html"] = "H4sIAAAAAAAA/+xa/4/bthX/PUD+hxftB9s4W7or1qFzrEuXdCm6rUuGJgGGw2GgJdpmQ5EqSflyud7/PpDUF0qmLCs5LEM3/eCzyPceyc/7Tt/qyXevXrz55+s/w05l9PLxo5X+CxSxbRxgFpgRjNLLx48AAFYZVgiSHRISqzh4++bl4pugmqOEvQd1m+M4UPiDihIpAxCYxoFUtxTLHcYqgJ3AmziIooLl77dhwrNozbmSSqA8Sok0bM1QmBEWakHRg62y2Bf4W4oUlsou2Jo5XFARRfHlhlAMa5yuIvv++NEqKoFZrXl6q/+mZA8JRVLGQcKZQoRhodEpBa0XhOWFWmwFL3IgaRxQviXsJRdZhWFJt+Eis8TlSXMk5Q0XaQDLnKIE7zhNsYgDIv+mRTybGEk4nSwnir/HbBLAfpHxVMNiBoLLVeTKba1XKMVZtfO1YrBWbMELRQnDi1yQDInboNyJJQ7g24SS5H15hODS/FlFdrZCLmqduAFiJ2pw7dfyRW/P4FLklKPUA0xbHDhPhZrRU312/eYcXb+6Ag82OGopV0HaFANo6UZyoRw1mNd+LXy+JkrQzInt944+ho5ca+WhECgEdQDQb0es8Deoi7f6xA+gCieuKLSmeCGwzDmTZI+D7rYNQYu6S2LJnKDenhCe0ZIFZML1aRNOg0uGMryK1O5Eckk+jiFPeMHUCPos/XoEdYopVn3bWUVeDDR1H2ZlAjicELDXJhYHUx16fmAbPgfCUvxhBoRBNSZ9KrIC0ssVAoXEVqfbf60pYu8DWNrsVrGHxrPu7up3rZr7+1WELleR8u24ku3w6C8/kY+GbwzPC62oQaZVwlPssmbp15rJDA/w+mdg2EdTxLZYdFy0X1r17BfEBZdIE1Jr17a2o8emVpN9yqsey3DkFIehoT3txaDfSj3GuIpMJGgFn5Ts6zyto+Bwej4Whkza3pho+6YbdFbrhT1hXwQv1bhHgiCm4uAGCUbYtsE8oRiJP1H6AiU7HFyaV0CUQqIHdCD1RtfOqs4kQ3vQqy0oWuuEsBYYpYkosvVBQOW0tq9eIrCF6SHhgiicBXUQaCbcMNCMyh7zrHBoKL/ntfWtUFnt/oz2SCaC5Gr5VMeDhroVESg5sA5OW8gxtG/DNZxUvAnFn0xGJJIRSWRMAjk9eZycODzu6E8Y3mTxWYmiTBIeG6jthueYtSLWmGQxNlGMThKfmiD6k8PDJ4aHTArHEsKRZOBBwW91HQNrB/8q8Ov+1X5bWXsBKZJ2x9ztk6vu+Gepa2fLdYx/TMd9TOZOqVwuoyhJWag5EymN/F9k9Ifwm/A8+kWesK8uN/pAuIzOw4s/hueLNVYovLBjPlm1tBLFG8JSfhNytsYbLnDBdI0PMUzxHjM1g/gSJhm6XWOQPMOQIoWAcQUS7fGkCq0JZ1KB6ZzfFRhiYPgG3hV4etcoD9MlTH5X3xNM5s2UFrqEu7b6TbO/BFZQOm/PlFcFS9ggKh3ru3foMqx2PJUHUmnJWbBEEc5gOutS6Meil3OppkFk7wXm8A8ZSiUI25LN7fSu3J7aERma7/ezmd8RQrXDbFpWGFgD6lmxeiySAsuCKoih4go1RP1cpRIRxUJNLXOYISnRFvdsyuC4gYpYByeIY7jwwuE+zYG1mgtKTyAvFQYxKFEcqR/1YyuvdwUOKZFKx6FX4qUZ0+FqeuQ494CpxCftv9lQx4QOZPqn7vtUnSCV7KZYCC60olt6CczwEgI4g7/89OrvjjWZmVnXgJzFy696XdfjmouSAbdrrqEG/U5yoZYQBB2n0wnD642FoN7xOtMv4ep6Ps5Lm3MNuqo2YqNTvW9twnorvVbsUELgSZQefVukcyRQViKsYfwOKeQzRkMXojzHLJ3a25l5s+ogg7nlKxn0dw9DKzahNCMscu7L5lbgHO50xYaFxnbygjOFmVq8uc3xZAmTrKCK5Ego0y4ttBlM7ntt+rcXvo4YwAGtuYU9JdKVW9EnrStLAwFK0ZqSj/j5rcJy6iM7csjWPow7hQWTO7JRrqRjUfHzItinhq5ukKhuEb+8P/f6z1tBD7K8DYb1KnMb7cx7Iej/Usof4zOFoP93mc91maYRc32m7KQ9GnMat/bpq/a7f7me6qIuxY5VFs5N2WBpkRRCYKZeI7XzVBjO5dESrg4PeKf7+yUEgnNl8pyV0sXt+gHrEPeW6jQtGOi3XHFbs1pVOCe7MqzXod79gAFUVx2nLbxHoj6qLrRbJlCu6o+2nauA3nBQdYo5ZlP3wr7HLQRWhWAnheQuaLX0E1CSWD2v4R3MLlUxp3ayAskxylDmlOi0EHjO1NUkxHDVZ5IepClWsEFqh0Wz3R+0ViCG80PyDRcw1TxEzz8FAiu77ZBitlW7p0DOznpVpdVqqK/IdZVD4ddfwR0LgqOBP+FMEdbXrfWEyC5GYV7I3fTIKha/alvzfkIL7YE3eRG13gVnEEQ6+lbSx0V1r+Szs0F77tqnr4s98f5hi5sSxScmmMOdqfZ14No4k68bsBzrHqjwmy1VFciQfXxC1fJJ9YfDZK6fHIN2SwPrGHr2fFAkdCuoQKMFOyTN5ZYp+rkoM6AvGnSf1v7gqsfYqqfHfaqnuQOpI7iFuFpiNLsp36ZTwjb8QtcDG/7VbKAsrR6NvuEr8wM8ia2AoXzRfWxCgCctYc9gcQFLuBgWMYCZs4JdwMSARblV/TKgkgEl+wKyD+hTwnNb7oGMK3J9rBwe4Bg4xxEYT7tA+7TG47+jpm5qDDcEmyqjt5hzIijE4Dck2yG5lYjvbshQnXKr+QVbAf20fqM+qWF/UhpFwtmGiGwayEJg6Py4Dc+ePLO669alts7rdZcRlaSnuW//4t5t8L9kC/8f691H36R/cWc96EhrI3Rs+8CsXTCMWR4YWk/KapcCVQXQzftdO7zvrOZ1ArtfW1GMdYAjC3oMXeCM73F5C9y2cr2eLQq9Kx84wUnGP87oTzT2UUbulDc5JUn5i/UcLjyCu2V698SfZ8P3B3bazdpr/dk6ie7YTRc5eT6Zw+Sv5vNH8/m9+XxjPl8/n1w/bXPpquBH3a9uKOdiar5Svq3WiKAeuTj/6vezmcNelkiWtKLM+Y2hnLtS8Ww2CxV/ST7gdPrVTLdTBgd5hav93D9+5P5KHVX/PR6Z/8D/dwAAAP//n4zsKJEvAAA="
}
