// Libs
use axum::{middleware::from_fn, routing::post, Router};
use middlewares::log::log_middleware;
use std::{net::SocketAddr, process::exit};
use tracing::error;

mod controllers;
mod middlewares;
mod serializations;
mod utils;

// Functions
pub fn get_app() -> Router {
    Router::new()
        .route("/encode", post(controllers::encode::post_encode))
        .layer(from_fn(log_middleware))
}

#[tokio::main]
async fn main() {
    // Initialize the logger.
    tracing_subscriber::fmt::init();

    // Listen to the port.
    let listener = match tokio::net::TcpListener::bind("0.0.0.0:8080").await {
        Ok(listener) => listener,
        Err(e) => {
            error!("Couldn\'t listen to port 8080. {}", e);
            exit(1);
        }
    };

    // Start the server.
    let app = get_app();
    if let Err(e) = axum::serve(
        listener,
        app.into_make_service_with_connect_info::<SocketAddr>(),
    )
    .await
    {
        error!("Couldn\'t start the server. {}", e);
        exit(1);
    }
}
