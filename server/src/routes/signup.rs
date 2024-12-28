use actix_web::{web, HttpResponse, Responder};
use serde::Deserialize;
use dotenv::dotenv;
use std::env;
use unofficial_appwrite::client::{Client, ClientBuilder};
use unofficial_appwrite::error::Error;
use unofficial_appwrite::services::server::users::Users;
use unofficial_appwrite::id::ID;

#[derive(Deserialize)]
pub struct SignupData {
    email: String,
    username: String,
    password: String
}

fn build_client(endpoint: &str, id: &str, key: &str) -> Result<Client, Error> {
    ClientBuilder::default()
    .set_endpoint(endpoint)?
    .set_project(id)?
    .set_key(key)?
    .build()
}

pub async fn signup(data: web::Json<SignupData>) -> impl Responder {
    dotenv().ok();

    let api_endpoint = match env::var("APPWRITE_API_ENDPOINT") {
        Ok(endpoint) => endpoint,
        Err(_) => return HttpResponse::InternalServerError().body("APPWRITE_API_ENDPOINT was not defined in the .env file")
    };
    
    let project_id = match env::var("APPWRITE_PROJECT_ID") {
        Ok(id) => id,
        Err(_) => return HttpResponse::InternalServerError().body("APPWRITE_PROJECT_ID was not defined in the .env file")
    };

    let api_key = match env::var("APPWRITE_API_KEY") {
        Ok(endpoint) => endpoint,
        Err(_) => return HttpResponse::InternalServerError().body("APPWRITE_API_ENDPOINT was not defined in the .env file")
    };
    
    let client= match build_client(&api_endpoint, &project_id, &api_key) {
        Ok(c) => c,
        Err(_) => return HttpResponse::InternalServerError().body("Client failed to build")
    };
    
    let create_user = match Users::create(
        &client, 
        ID::unique(), 
        Some(&data.email),
        None, 
        Some(&data.password), 
        Some(&data.username)
    ).await {
        Ok(c) => c,
        Err(e) => return HttpResponse::InternalServerError().body(format!("{:?}", e))
    };
    
    HttpResponse::Ok().body(format!("{} has signed up successfully!", create_user.name))
}