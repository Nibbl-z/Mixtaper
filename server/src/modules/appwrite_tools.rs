use dotenv::dotenv;
use serde_json::json;
use unofficial_appwrite::models::user::User;
use unofficial_appwrite::query::Query;
use unofficial_appwrite::services::server::account::Account;
use unofficial_appwrite::services::server::users::Users;
use std::env;

use unofficial_appwrite::client::{Client, ClientBuilder};
use unofficial_appwrite::error::Error;

fn build_client(endpoint: &str, id: &str, key: &str) -> Result<Client, Error> {
    ClientBuilder::default()
    .set_endpoint(endpoint)?
    .set_project(id)?
    .set_key(key)?
    .build()
}

fn build_client_with_token(endpoint: &str, id: &str, key: &str, token: &str) -> Result<Client, Error> {
    ClientBuilder::default()
    .set_endpoint(endpoint)?
    .set_project(id)?
    .set_key(key)?
    .set_jwt(token)?
    .build()
}

pub fn get_client(token: Option<&str>) -> Result<Client, String> {
    dotenv().ok();
    
    let api_endpoint = match env::var("APPWRITE_API_ENDPOINT") {
        Ok(endpoint) => endpoint,
        Err(_) => return Err(String::from("APPWRITE_API_ENDPOINT was not defined in the .env file"))
    };
    
    let project_id = match env::var("APPWRITE_PROJECT_ID") {
        Ok(id) => id,
        Err(_) => return Err(String::from("APPWRITE_PROJECT_ID was not defined in the .env file"))
    };
    
    let api_key = match env::var("APPWRITE_API_KEY") {
        Ok(endpoint) => endpoint,
        Err(_) => return Err(String::from("APPWRITE_API_KEY was not defined in the .env file"))
    };

    match token {
        None => match build_client(&api_endpoint, &project_id, &api_key) {
            Ok(c) => return Ok(c),
            Err(_) => return Err(String::from("Client failed to build"))
        },
        Some(t) => match build_client_with_token(&api_endpoint, &project_id, &api_key, &t) {
            Ok(c) => return Ok(c),
            Err(_) => return Err(String::from("Client failed to build"))
        },
    }
}

pub async fn get_user(token: &str) -> Result<User, String> {
    let client = get_client(Some(token))?;

    match Account::get(&client).await {
        Ok(user) => Ok(user),
        Err(error) => Err(error.to_string())
    }
}

pub async fn get_user_by_username(client: &Client, username: &str) -> Result<User, String> {
    match Users::list(&client, Some(vec![Query::equal("name", json![vec![username]])]), None).await {
        Ok(list) => Ok(list.users[0].clone()),
        Err(error) => Err(error.to_string())
    }
}