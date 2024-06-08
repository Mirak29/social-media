import { makeRequest } from '$lib/api'

import { DB } from '../../routes/db'

export const authenticateUser =  async (Cookies, locals) => {
    let resp=await makeRequest("","GET",{},{},Cookies)
    if (locals!=undefined && !locals?.user) {
        locals.user=resp?.data?.user
    }
    console.log("local data is her",locals);
    DB("set","user",resp?.data?.user)

    return resp?.data?.isauth
}
