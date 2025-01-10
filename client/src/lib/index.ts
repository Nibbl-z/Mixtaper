// place files you want to import through the `$lib` alias in this folder.

import type { MessageResult } from "./Types";

export function getCookie(name: string) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);

    if (parts.length === 2) {
        const cookieValue = parts.pop();
        if (cookieValue) {
            return cookieValue.split(';').shift();
        }
    }
}

export function deleteCookie(name: string) {
    document.cookie = name + '=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}

export async function getID(): Promise<string> {
    const token = getCookie("token");
    
    if (token === undefined) {return ""}
    
    const response = await fetch(`http://localhost:2050/get_my_id`, {
        method: "GET",
        headers: {
            "Authorization": token
        }
    })
    
    const data: MessageResult = await response.json()
    if (data.successful) {
        return data.message
    } else {
        return ""
    }
}

export async function getPfp(): Promise<string> {
    const id = await getID()
    return (id == "") ? "/account.png" : `https://cloud.appwrite.io/v1/storage/buckets/profile_pictures/files/${id}/view?project=676f205d000370a15786&project=676f205d000370a15786&mode=admin?` + new Date().getTime();
}