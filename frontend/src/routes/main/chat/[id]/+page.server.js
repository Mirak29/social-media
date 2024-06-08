import { makeRequest } from '$lib/api.js'
import { authenticateUser } from '$lib/auth/auth.js';
import { error, redirect } from '@sveltejs/kit';

export const load = async ({ cookies, params ,locals}) => {
    const IsAuth = await authenticateUser(cookies,locals)
    if (!IsAuth) {
        redirect(302, "/login")
    }
    let receiverid = params.id;
    let response = await makeRequest(`chat?receiver=${receiverid}`, "GET", {}, {}, cookies);
    if (response.status === 200) {
        let res = {
            result: response.data
        }
        return{ res }
    } else {
        throw error(403, "Not Allowed");
    }
}
