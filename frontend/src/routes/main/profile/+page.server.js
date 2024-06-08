import { makeRequest } from "$lib/api.js"
import { authenticateUser } from "$lib/auth/auth"
import { error, redirect } from "@sveltejs/kit"

export const load = async ({ cookies,locals }) => {
    const IsAuth = await authenticateUser(cookies,locals)
    if (!IsAuth) {
        redirect(302, "/login")
    }
    const response = await makeRequest("profile", "get", {}, {}, cookies)

    if (response?.data.success) {
        return response?.data
    }
    throw error(500, "internal server error")
}