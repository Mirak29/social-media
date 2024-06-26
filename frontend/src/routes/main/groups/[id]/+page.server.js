import { makeRequest } from "$lib/api";
import { saveImage } from "../../../../lib/index.js";
import axios from 'axios';
import { localStorageObj } from "../../../db.js";


export const actions = {
    createEvent: async ({ request, cookies, params }) => {
        const formDatas = await request.formData();
        const groupId = params.id;
        formDatas.append('groupId', groupId);


        let title = formDatas.get('title');
        let description = formDatas.get('description');
        let date = formDatas.get('date');
        let time = formDatas.get('time');
         let errorMsg = '';
 
         if (title.length <=  2 || title.length > 30 || description.length <= 2 || title.length > 100) {
             errorMsg = 'Invalid input entered';
         }
 
         const today = new Date();
         const eventDate = new Date(date);
         if (eventDate <= today) {
             errorMsg = 'Invalid Date';
         }
 
         const timeRegex = /^([01]?[0-9]|2[0-3]):[0-5][0-9]$/;
         if (!timeRegex.test(time)) {
             errorMsg = 'invalid Time (format HH:MM).';
         }

        if (errorMsg.length < 1) {
            let response = await makeRequest("createEvent", "POST", formDatas, {}, cookies)
            if (response.data.success == false) {
                errorMsg = response.data.error
            } else {
                errorMsg = 'no'
                title = ''
                description = ''
                date = ''
                time = ''
            }
        }

        let res = {
            error: errorMsg,
            title: title,
            description: description,
            date: date,
            time: time,
        }

        return res;
    },
    createPost: async ({ request, cookies, params }) => {
        // TODO log the user in
        let data = await request.formData()
        let content = data.get('content')
        if (content == "") {
            return fail(400, { content, missing: true })
        }
        let post = {
            titre: "",
            content: content,
            image: await saveImage(data.get("avatar")),
            privacy: "groupe",
            Group_id: Number(params.id)
        }
        let response = await makeRequest("addPost", "POST", post, {}, cookies)
    },
    createComment: async ({ request, cookies }) => {
        // TODO log the user in
        let data = await request.formData()
        let content = data.get("comment")

        let postId = data.get("postId")
        let user = JSON.parse(localStorageObj?.data?.user)

        let comment = {
            postId: Number(postId),
            comment: content,
            image: await saveImage(data.get("avatar")),
            firstName: user.FirstName,
            lastName: user.LastName,
        }

        const response = await makeRequest("addComment", "POST", comment, {}, cookies)

        if (response.status == 200) {
            return { succes: true, data: comment }
        } else {
            return { success: false }
        }

    }

};

// export const load = async ({cookies, params}) => {



// };
export const load = async ({ cookies, params }) => {
    const groupId = params.id;
    const data = { groupId: groupId }
    let response = await makeRequest(`getgroupdetail/${groupId}`, "GET", data, {}, cookies);
    if (response.status === 200) {

        let res = {
            result: response.data
        }
        return { res }
    }
};