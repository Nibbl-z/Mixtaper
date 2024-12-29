use actix_web::{web, App, HttpServer};
use actix_cors::Cors;

mod routes {
    pub mod signup;
    pub mod login;

    pub mod upload_riq;
}

pub mod modules {
    pub mod appwrite_tools;
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(move || {
        App::new()
        .wrap(
            Cors::default()
            .allow_any_origin()
            .allow_any_header()
            .allow_any_method()
        )
        .service(web::resource("/signup").route(web::post().to(routes::signup::signup)))
        .service(web::resource("/login").route(web::post().to(routes::login::login)))
        .service(web::resource("/upload_riq").route(web::post().to(routes::upload_riq::upload_riq)))
    })
    
    .bind(("localhost", 2050))?.run().await
}