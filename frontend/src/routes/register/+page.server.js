import { makeRequest } from "$lib/api.js";
import { authenticateUser } from "$lib/auth/auth.js";
import { generateRandom, saveImage } from "$lib/index.js";

import { redirect } from "@sveltejs/kit";
import { localStorageObj } from "../db.js";

export const load = async ({ cookies,locals }) => {
    if (localStorageObj?.data?.user) {
        const IsAuth = await authenticateUser(cookies,locals)
        if (IsAuth) {
            redirect(302, "/")
        }
    }
}

const validateFormData = (formData) => {
    const emailRegex = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/;
    const fields = ['email', 'password', 'firstname', 'lastname'];

    for (let field of fields) {
        const value = formData.get(field);
        if (!value || value.length < 2 || value.length > 50) {
            return [false, field];
        }
        if (field === 'email' && !emailRegex.test(value)) {
            return [false, field];
        }
    }
    const value = formData.get("bio");
    if (value.length >150) {
        return [false, "bio"];
    }
    const value1 = formData.get("nickname");
    if (value1.length >50) {
        return [false, "nickname"];
    }
    return [true, ''];
}

export const actions = {
    default: async ({ request, }) => {
        const formDatas = await request.formData()
        let errorMsg = '';
        let [bool, err] = validateFormData(formDatas)
        if (bool == true) {
            let imagePath = '';
            let imageName = formDatas.get('avatar').name

            if (imageName != 'undefined' && imageName != "") {
                let image = await saveImage(formDatas.get('avatar'));
                if (image !== '') {
                    imagePath = image;
                    formDatas.set('avatar', imagePath);
                } else {
                    errorMsg = 'Invalid image size or format';
                }
            }
            if (errorMsg.length == 0) {
                let response = await makeRequest("register", "POST", formDatas)
                if (response?.data?.success == true) {
                    redirect(302, "/login")
                }
                errorMsg = response?.data?.error
            }
        } else {
            errorMsg = err.toUpperCase() + ' invalid. Veuillez verifier et reesayer.'
        }

        let res = {
            error: errorMsg,
            email: formDatas.get('email'),
            firstname: formDatas.get('firstname'),
            lastname: formDatas.get('lastname'),
            birth: formDatas.get('birthdate'),
            nickname: formDatas.get('nickname'),
            bio: formDatas.get('bio')

        }

        return res;

    }
};

