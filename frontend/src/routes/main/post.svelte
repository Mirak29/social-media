<script>
  import Comment from "./comments.svelte";
  import { makeRequest } from "$lib/api.js";
  import { PUBLIC_BACKEND_URL_FRONTEND,PUBLIC_SOCKET_URL } from '$env/static/public';
  import axios from "axios";

	export let data;
	
	function calculateAgo(dateString) {
		// Convert the date string to a Date object
		const date = new Date(dateString.split(".")[0]);
		const now = new Date();

		// Calculate the difference between the dates in milliseconds
		const difference = now - date;

		// Convert the difference to minutes, hours, or days
		const minutes = Math.floor(difference / (1000 * 60));
		const hours = Math.floor(difference / (1000 * 60 * 60));
		const days = Math.floor(difference / (1000 * 60 * 60 * 24));

		// Determine the age
		if (days > 0) {
			return days + (days > 1 ? " days ago" : " days ago");
		} else if (hours > 0) {
			return hours + (hours > 1 ? " hours ago" : " hour ago");
		} else {
			return minutes + (minutes > 1 ? " minutes ago" : " minute ago");
		}
	}

  async function getPostDetails(postId) {
    let url = `${PUBLIC_BACKEND_URL_FRONTEND}/getPost?postid=${postId}`;
    try {
           let header={
                cookie:document.cookie
            }
            const config = { method:"get",withCredentials: true , header,mode: 'no-cors' };
            let response= await axios(url,config)
           console.log("SUCCES POSTDETAIL",  response.data);

            return response.data
        } catch (err) {
           console.log("ERORR POSTDETAIL",  err);
        }
  }
  
  export let CommSection
  async function ShowComments(postId) {
    if (CommSection?.display == "none") {
      console.log("show comments ", postId)
      console.log("hello " , CommSection.data)
      CommSection.display = "block"
    } 
    CommSection.data = await getPostDetails(postId)
  }

</script>

{#if data}
  {#each data as post}
    <div class="card w-100 shadow-xss rounded-xxl border-0 p-4 mb-4">
      <div class="card-body p-0 d-flex">
        <figure class="avatar me-3">
          {#if !post.avatar}
              <img src="//ui-avatars.com/api/?name={post.firstName +' '+post.lastName}e&size=100&rounded=true&color=fff&background=random"
            alt="{post.firstName +' '+post.lastName}"
            class="shadow-sm rounded-circle w45"
          />
          {:else}
            <img src="{post.avatar}" alt="{post.firstName}" class="shadow-sm rounded-circle w45" style="height: 50px;"/>
          {/if}
          
        </figure>
        {#if post?.ispublic === 1}
          <a href="/main/profile/{post.User_id}">
            <h4 class="fw-700 text-grey-900 font-xssss mt-1">
              {post.firstName +' '+post.lastName} <span
                class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500"
                >{calculateAgo(post.creationDate)}</span
              >
            </h4>
          </a>
        {:else}
        <h4 class="fw-700 text-grey-900 font-xssss mt-1">
          {post.firstName +' '+post.lastName} <span
            class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500"
            >{calculateAgo(post.creationDate)}</span
          >
        </h4>
        {/if}

      </div>
      <div class="card-body p-0 me-lg-5">
        <p class="fw-500 text-grey-500 lh-26 font-xssss w-100 mb-2 text-break">
          {#if post?.content.length > 200}
                {#if post.display || post.display == undefined}
                      {post.content.slice(0,200)}
                      <a class="fw-600 text-primary ms-2" on:click={()=>{if(post.display == undefined ){post.display = false}else{post.display = !post.display}}} style="cursor:pointer;" >See more</a>
                {:else}
                      {post.content}
                      <a class="fw-600 text-primary ms-2" on:click={()=>{post.display = !post.display}} style="cursor:pointer;">Show Less</a>
                {/if}
          {:else}
                {post.content}
          {/if}
        </p>
      </div>
      <div class="card-body d-block p-0 mb-3">
        <div class="row ps-2 pe-2">
          <div class="col-sm-12 p-1" >
              <img 
                src={post.image}
                class="rounded-3 w-100"
                alt=""
              />
          </div>
        </div>
      </div>
      <div class="card-body d-flex p-0" on:click={()=> {ShowComments(post.id)}}>
        <a
          style="cursor: pointer;"
          class="d-flex align-items-center fw-600 text-grey-900 text-dark lh-26 font-xssss"
          ><i
            class="feather-message-circle text-dark text-grey-900 btn-round-sm font-lg"
          ></i><span class="d-none-xss">Comments</span></a
        >
      </div>
    </div>
  {/each}
{/if}
