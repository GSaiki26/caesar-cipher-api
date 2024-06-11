// Data
const ALPHABET: &'static str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";

// Functions
/**
A method that encodes a text to Caesar cipher.

## Arguments
- `text` - A string slice that holds the text to be encoded.
- `shift` - An i8 that holds the shift value.
*/
pub fn encode_to_caesar_cipher(text: &str, shift: usize) -> String {
    let mut encoded_text = String::new();
    for char in text.chars() {
        // Find the char in the alphabet.
        let is_upper = char == char.to_ascii_uppercase();
        let index = match ALPHABET.find(char.to_ascii_uppercase()) {
            Some(index) => (index + shift as usize) % ALPHABET.len(),
            None => {
                encoded_text.push(char);
                continue;
            }
        };

        // Replace the selected char by the char with the given shift.
        let mut char = ALPHABET.chars().nth(index).unwrap();
        if !is_upper {
            char = char.to_ascii_lowercase();
        }
        println!("{char}");
        encoded_text.push(char);
    }

    encoded_text
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_encode_to_caesar_cipher() {
        let text = "Hello world!";
        assert_eq!(encode_to_caesar_cipher(text, 0), text);
        assert_eq!(encode_to_caesar_cipher(text, 25), "Gdkkn vnqkc!");
    }
}
