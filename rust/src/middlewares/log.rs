// Libs
use std::net::SocketAddr;

use axum::{
    body::Body,
    extract::{self, Request},
    middleware::Next,
    response::Response,
};
use tracing::{info, info_span, Instrument};

// Functions
/**
The log middleware. It'll be responsible for defining the spam.
*/
pub async fn log_middleware(
    connect_info: extract::ConnectInfo<SocketAddr>,
    req: Request,
    next: Next,
) -> Response<Body> {
    let span = info_span!("", ip = connect_info.to_string());
    let _guard = span.enter();

    info!("Request received: {} {}", req.method(), req.uri());
    let res = next.run(req).in_current_span().await;
    info!("Response returned: {}", res.status());
    return res;
}
