use actix_web::{web, HttpResponse, Responder};
use serde::Deserialize;
use bcrypt::{DEFAULT_COST, hash};

use unofficial_appwrite::services::server::users::Users;
use unofficial_appwrite::id::ID;

use crate::modules::appwrite_tools;

#[derive(Deserialize)]
pub struct SignupData {
    email: String,
    username: String,
    password: String
}

pub async fn signup(data: web::Json<SignupData>) -> impl Responder {
    let client = match appwrite_tools::get_client(None) {
        Ok(c) => c,
        Err(error) => return HttpResponse::InternalServerError().body(error)
    };
    
    let hashed_password = match hash(&data.password, DEFAULT_COST) {
        Ok(pass) => pass,
        Err(_) => return HttpResponse::InternalServerError().body("Password failed to hash")
    };
    
    let create_user = match Users::create_bcrypt_user(
        &client, 
        ID::unique(), 
        &data.email,
        &hashed_password, 
        Some(&data.username)
    ).await {
        Ok(c) => c,
        Err(e) => return HttpResponse::InternalServerError().body(format!("{:?}", e))
    };
    
    HttpResponse::Ok().body(format!("{} has signed up successfully!", create_user.name))
}