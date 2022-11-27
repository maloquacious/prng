# PRNG
PRNG provides simple random number generators for testing.

Usage:

    // create a new LCG generator
    lcg32 := prng.LCG32(0)

    // roll 5 6-sided dice
    r1 := sfc32.Roll(5, 6)

    // create a new SFC generator
    sfc32 := prng.SFC32(0, 12345, 0, 1)

    // roll 5 6-sided dice
    r2 := sfc32.Roll(5, 6)

## Features
- 32-bit generators
- LCG and SFC implementations

## Installation
No installation required - just import the package into your program.

    import github.com/mdhender/prng

# Contribute
- Issue Tracker: github.com/mdhender/prng/issues
- Source Code: github.com/mdhender/prng

## Support
If you are having issues, please let us know.

## License
The project is licensed under the MIT license.

## Thanks
Special thanks to Eric Holscher and [Write The Docs](https://www.writethedocs.org/guide/writing/beginners-guide-to-docs/).

### LCG32
Constants for LCG32 from Open Adventure.

### SFC32
Code for SFC32 is from [SimBlob](https://simblob.blogspot.com/2022/05/upgrading-prng.html#more).
