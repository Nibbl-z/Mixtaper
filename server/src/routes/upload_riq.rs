use std::{fs::File, io::Write, fs::remove_file};

use actix_web::{web::Bytes, HttpRequest, HttpResponse, Responder};

use unofficial_appwrite::models::user;
use unofficial_appwrite::{id::ID, services::server::storage::Storage};
use unofficial_appwrite::permission::Permission;
use unofficial_appwrite::role::Role;
use crate::modules::appwrite_tools;

pub async fn upload_riq(req: HttpRequest, body: Bytes) -> impl Responder {
    /*let session_token = match req.headers().get("Authorization") {
        Some(value) => {
            let token = value.to_str();
            if let Err(error) = token { return HttpResponse::InternalServerError().body("Failed to get token"); }

            token.unwrap()
        },
        None => return HttpResponse::Unauthorized().body("Authorization token is missing")
    };

    let user = match appwrite_tools::get_user(session_token).await {
        Ok(c) => c,
        Err(error) => return HttpResponse::InternalServerError().body(error)
    };*/

    let directory = format!("uploads/{}.riq", "test");

    let client = match appwrite_tools::get_client(None) {
        Ok(c) => c,
        Err(error) => return HttpResponse::InternalServerError().body(error)
    };

    let mut file = match File::create(&directory) {
        Ok(f) => f,
        Err(error) => return HttpResponse::InternalServerError().body(format!("Failed to create uploaded file: {}", error.to_string()))
    };

    if let Err(error) = file.write(body.as_ref()) {
        return HttpResponse::InternalServerError().body(format!("Failed to write to uploaded file: {}", error.to_string()))
    }

    let uploaded_file = match Storage::create_files(
        &client, 
        "riq_files", 
        ID::unique(), 
        &directory, 
        String::from("test.riq"), 
        None).await {
            Ok(f) => f,
            Err(error) => return HttpResponse::InternalServerError().body(format!("Failed to upload file to server: {}", error.to_string()))
        };

    if let Err(error) = remove_file(&directory) {
        return HttpResponse::InternalServerError().body(format!("Temp file failed to delete: {}", error.to_string()))
    }

    HttpResponse::Ok().body(uploaded_file.id)

    /* Some(vec![
        Permission::read(&Role::user(&user.id, None)),
        Permission::update(&Role::user(&user.id, None)),
        Permission::delete(&Role::user(&user.id, None))
    ]) just gotta move this to other pc and fix authentication */
}