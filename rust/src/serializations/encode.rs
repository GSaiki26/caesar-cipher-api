// Libs
use serde::{Deserialize, Serialize};

// Structs
#[derive(Debug, Deserialize, Serialize)]
pub struct ReqEncode {
    pub content: String,
}

#[derive(Debug, Default, Deserialize, Serialize)]
pub struct ResEncode {
    pub results: Vec<ResEncodeResult>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct ResEncodeResult {
    pub shift: i8,
    pub text: String,
}
