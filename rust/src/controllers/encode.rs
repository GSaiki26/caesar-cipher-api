// Libs
use crate::{
    serializations::encode::{ReqEncode, ResEncode, ResEncodeResult},
    utils::utils::encode_to_caesar_cipher,
};
use axum::{http::StatusCode, response::IntoResponse, Json};
use tracing::info;

// Functions
/**
POST /encode
*/
pub async fn post_encode(Json(body): Json<ReqEncode>) -> impl IntoResponse {
    info!("Encoding provided content...");

    let mut res = ResEncode::default();
    for shift in 0..=25 {
        res.results.push(ResEncodeResult {
            shift,
            text: encode_to_caesar_cipher(&body.content, shift as usize),
        });
    }

    info!("Encoding completed successfully");
    (StatusCode::OK, Json(res))
}
