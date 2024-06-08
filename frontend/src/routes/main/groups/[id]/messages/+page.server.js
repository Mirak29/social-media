import { makeRequest } from '$lib/api.js'
import { error } from '@sveltejs/kit';

export const load = async ({cookies , params})=>{
    let groupid = params.id;
    let response = await makeRequest(`chatgroup?groupid=${groupid}`, "GET", {}, {}, cookies);
    if (response.status ===  200) {
        
        let res = {
            result : response.data
        }
        return{ res }
    }else {
        throw error(403, "Not Allowed");
    }
}
