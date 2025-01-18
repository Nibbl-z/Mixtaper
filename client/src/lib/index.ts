// place files you want to import through the `$lib` alias in this folder.

import { PUBLIC_BACKEND_URL, PUBLIC_PROJECT_ID } from "$env/static/public";
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
    
    const response = await fetch(`${PUBLIC_BACKEND_URL}/get_my_id`, {
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

export async function getPfpFromId(userId: string, placeholder: string | "/placeholder_pfp.png"): Promise<string> {
    const response = await fetch(`https://cloud.appwrite.io/v1/storage/buckets/profile_pictures/files/${userId}/view?project=${PUBLIC_PROJECT_ID}&project=${PUBLIC_PROJECT_ID}&mode=admin`)
    
    if (response.ok) {
        return `https://cloud.appwrite.io/v1/storage/buckets/profile_pictures/files/${userId}/view?project=${PUBLIC_PROJECT_ID}&project=${PUBLIC_PROJECT_ID}&mode=admin?` + new Date().getTime();
    } else {
        return placeholder
    }
}

export async function getPfp(placeholder: string): Promise<string> {
    const id = await getID()
    return (id == "") ? "/account.png" : getPfpFromId(id, placeholder)
}