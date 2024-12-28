use actix_web::{web, HttpResponse, Responder};
use serde::Deserialize;
use unofficial_appwrite::services::server::account::Account;

use crate::modules::appwrite_tools;

#[derive(Deserialize)]
pub struct LoginData {
    identifier: String,
    password: String
}

pub async fn login(data: web::Json<LoginData>) -> impl Responder {
    let client = match appwrite_tools::get_client(None) {
        Ok(c) => c,
        Err(error) => return HttpResponse::InternalServerError().body(error)
    };
    
    let email = if data.identifier.contains("@") { 
        &data.identifier
    } else { // Not an email!
        match appwrite_tools::get_user_by_username(&client, &data.identifier).await {
            Ok(user) => &user.email.clone(),
            Err(error) => return HttpResponse::InternalServerError().body(error)
        }
    };

    let session = match Account::create_email_password_session(&client, &email, &data.password).await {
        Ok(session) => session,
        Err(error) => return HttpResponse::Unauthorized().body(error.to_string())
    };
    
    HttpResponse::Ok().body(session.secret)
}