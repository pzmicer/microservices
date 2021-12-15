#[macro_use] extern crate rocket;

#[get("/")]
fn index() -> String {
    let name = hostname::get().unwrap().into_string().unwrap();
    format!("Hello from Rocket!\nHostname: {}", name)
}

#[launch]
fn rocket() -> _ {
    rocket::build()
        .mount("/", routes![index])
}
