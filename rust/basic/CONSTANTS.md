# Constants 

- [Constants](https://doc.rust-lang.org/rust-by-example/custom_types/constants.html)

```rust
mod network;
use network::server::connect;

mod basic;
use basic::constants;

fn main() {
    constants::basic_constants();
}
```

```rust
pub mod constants;
```

```rust
pub mod constants {
    static LANGUAGE: &str = "Rust";
    const THRESHOLD: i32 = 10;

    fn is_big(n: i32) -> bool {
        n > THRESHOLD
    }
    pub fn basic_constants() {
        let n = 16;

        println!("This is {}", LANGUAGE);
        println!("The threshold is {}", THRESHOLD);
        println!("{} is {}", n, if is_big(n) {"big"} else { "small"})
    }
}
```