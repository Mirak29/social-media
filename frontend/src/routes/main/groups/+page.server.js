import { makeRequest } from "$lib/api";
import { authenticateUser } from "$lib/auth/auth.js";
import { redirect } from "@sveltejs/kit";

export const load = async ({cookies,locals}) => {
    const IsAuth=await authenticateUser(cookies,locals)
    if (!IsAuth) {
        redirect(302,"/login")
    }
    let response = await makeRequest("groups", "GET", {}, {}, cookies);
    if (response.status ===  200) {
        
        let res = {
            result : response.data
        }
        return{ res }
    }
};

export const actions = {
    invitegroup: async ({request,cookies}) => {
          const formDatas= await request.formData()
            let data={
                groupId:formDatas.get("groupId"),
                userId:formDatas.get("userId"),
                userIds : formDatas.getAll("usersSelect").map((v) => Number(v))
           }
            try {
                const response = await makeRequest("invitegroup", "POST", data, {}, cookies)
                console.log("RESPONSE TO INVITATION" , response )
                return {result: true, formdata: data}
            } catch (error) {
                 console.error('Failed to create group:', error);
                 return {result: true, formdata: data}
          }
      },
}
