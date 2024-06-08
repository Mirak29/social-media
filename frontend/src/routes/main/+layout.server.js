import { makeRequest } from "$lib/api.js"
import { authenticateUser } from "$lib/auth/auth"

import { error, redirect } from "@sveltejs/kit"

export const load= async ({cookies,locals})=>{
    const IsAuth=await authenticateUser(cookies,locals)
    if (!IsAuth) {
        redirect(302,"/login")
    }

    const response= await makeRequest("notificatons","get",{},{},cookies)
    if (response?.data?.success) {
        console.log("local data is her side bar",locals);
        return{notif: response?.data, user:locals?.user}
    }

    throw error(400,"bad request")
}