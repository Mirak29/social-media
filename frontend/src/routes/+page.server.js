import { makeRequest } from "$lib/api.js";
import { authenticateUser } from "$lib/auth/auth.js";
import { redirect } from "@sveltejs/kit";


export const load = async ({cookies,locals})=>{
    const IsAuth=await authenticateUser(cookies,locals)
    if (!IsAuth) {
        redirect(302,"/login")
    }
    redirect(302,"/main")
}