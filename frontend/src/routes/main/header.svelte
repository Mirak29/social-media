<script>
    
    export let data    
    import {onMount} from "svelte"
    import {WS } from "./socket"
    import {darkMode} from "./stores"
    import Contacts from "./messagerie.svelte";
    import { PUBLIC_SOCKET_URL } from '$env/static/public';
    let Users = []
    
// let logut=async (cookies)=>{
    import { displayContacts } from "./stores";
    import Notification from "./notification.svelte";
    const toggleConstants = () =>{
        displayContacts.update((show) => !show)
    }
	onMount(async () => {
		let ws = new WebSocket(PUBLIC_SOCKET_URL+"/ws")
		WS.set(ws)
        
	})
    let active = false
    let activeSidebar = false



</script>

    <div class="nav-header bg-white shadow-xs border-0">
        <div class="nav-top">
            <a href="/main"><i class="feather-zap text-success display1-size me-2 ms-0"></i><span class="d-inline-block fredoka-font ls-3 fw-600 text-current font-xxl logo-text mb-0">BuzzHub. </span> </a>
            <a href="#" class="mob-menu ms-auto me-2 chat-active-btn"><i class="feather-message-circle text--900 font-sm btn-round-md bg-greylight font-xl text-current" on:click={toggleConstants}></i></a>
            <!-- notif -->
            <!-- <a href="#" class="mob-menu me-2" id="dropdownMenu3" data-bs-toggle="dropdown" aria-expanded="false"><i class="feather-bell text-grey-900 font-sm btn-round-md bg-greylight"></i></a> -->
            <a href="/main" class="mob-menu me-2"><Notification {data}/></a>
            <!-- <a href="/main/logout" class="p-2 text-center ms-3 menu-icon "><i class="btn-round-md bg-red-gradiant text-white feather-lock font-md me-3"></i></a> -->
            <a href="/main/logout" class="me-2 menu-search-icon mob-menu"><i class="btn-round-md bg-red-gradiant text-white feather feather-log-out font-md me-3"></i></a>
            <button class="nav-menu me-0 ms-2 "  class:active={activeSidebar} on:click={()=> {activeSidebar = !activeSidebar
            document.querySelector(".navigation").classList.toggle("nav-active")
            }}></button>
        </div>
        <Notification {data}/>
        
        <a href="#" class="p-2 text-center ms-3 menu-icon chat-active-btn"  on:click={toggleConstants}><i class="feather-message-square font-xl text-current"></i></a>
        <div class="p-2 text-center ms-3 position-relative dropdown-menu-icon menu-icon cursor-pointer">
            <i class="feather-settings animation-spin d-inline-block font-xl text-current" on:click={()=> active = !active}></i>
            <div class="dropdown-menu-settings switchcolor-wrap " class:active={active}>
                <h4 class="fw-700 font-sm mb-4">Settings</h4>
                <h6 class="font-xssss text-grey-500 fw-700 mb-3 d-block">Choose Color Theme</h6>
                <ul>
                    <li>
                        <label class="item-radio item-content">
                            <input type="radio" name="color-radio" value="red" checked=""><i class="ti-check"></i>
                            <span class="circle-color bg-red" style="background-color: #ff3b30;"></span>
                        </label>
                    </li>
                    <li>
                        <label class="item-radio item-content">
                            <input type="radio" name="color-radio" value="green"><i class="ti-check"></i>
                            <span class="circle-color bg-green" style="background-color: #4cd964;"></span>
                        </label>
                    </li>
                    <li>
                        <label class="item-radio item-content">
                            <input type="radio" name="color-radio" value="blue" checked=""><i class="ti-check"></i>
                            <span class="circle-color bg-blue" style="background-color: #132977;"></span>
                        </label>
                    </li>
                    <li>
                        <label class="item-radio item-content">
                            <input type="radio" name="color-radio" value="pink"><i class="ti-check"></i>
                            <span class="circle-color bg-pink" style="background-color: #ff2d55;"></span>
                        </label>
                    </li>
                    <li>
                        <label class="item-radio item-content">
                            <input type="radio" name="color-radio" value="yellow"><i class="ti-check"></i>
                            <span class="circle-color bg-yellow" style="background-color: #ffcc00;"></span>
                        </label>
                    </li>
                    <li>
                        <label class="item-radio item-content">
                            <input type="radio" name="color-radio" value="orange"><i class="ti-check"></i>
                            <span class="circle-color bg-orange" style="background-color: #ff9500;"></span>
                        </label>
                    </li>
                    <li>
                        <label class="item-radio item-content">
                            <input type="radio" name="color-radio" value="gray"><i class="ti-check"></i>
                            <span class="circle-color bg-gray" style="background-color: #8e8e93;"></span>
                        </label>
                    </li>

                    <li>
                        <label class="item-radio item-content">
                            <input type="radio" name="color-radio" value="brown"><i class="ti-check"></i>
                            <span class="circle-color bg-brown" style="background-color: #D2691E;"></span>
                        </label>
                    </li>
                    <li>
                        <label class="item-radio item-content">
                            <input type="radio" name="color-radio" value="darkgreen"><i class="ti-check"></i>
                            <span class="circle-color bg-darkgreen" style="background-color: #228B22;"></span>
                        </label>
                    </li>
                    <li>
                        <label class="item-radio item-content">
                            <input type="radio" name="color-radio" value="deeppink"><i class="ti-check"></i>
                            <span class="circle-color bg-deeppink" style="background-color: #FFC0CB;"></span>
                        </label>
                    </li>
                    <li>
                        <label class="item-radio item-content">
                            <input type="radio" name="color-radio" value="cadetblue"><i class="ti-check"></i>
                            <span class="circle-color bg-cadetblue" style="background-color: #5f9ea0;"></span>
                        </label>
                    </li>
                    <li>
                        <label class="item-radio item-content">
                            <input type="radio" name="color-radio" value="darkorchid"><i class="ti-check"></i>
                            <span class="circle-color bg-darkorchid" style="background-color: #9932cc;"></span>
                        </label>
                    </li>
                </ul>
                
                <div class="card bg-transparent-card border-0 d-block mt-3">
                    <h4 class="d-inline font-xssss mont-font fw-700">Header Background</h4>
                    <div class="d-inline float-right mt-1">
                        <label class="toggle toggle-menu-color"><input type="checkbox"><span class="toggle-icon"></span></label>
                    </div>
                </div>
                <div class="card bg-transparent-card border-0 d-block mt-3">
                    <h4 class="d-inline font-xssss mont-font fw-700">Menu Position</h4>
                    <div class="d-inline float-right mt-1">
                        <label class="toggle toggle-menu"><input type="checkbox"><span class="toggle-icon" ></span></label>
                    </div>
                </div>
                <div class="card bg-transparent-card border-0 d-block mt-3">
                    <h4 class="d-inline font-xssss mont-font fw-700" >Dark Mode</h4>
                    <div class="d-inline float-right mt-1">
                        <label class="toggle toggle-dark"><input type="checkbox"><span class="toggle-icon" on:click={()=> {
                            document.body.classList.toggle("color-theme-blue")
                            document.body.classList.toggle("mont-font")
                            document.body.classList.toggle("loaded")
                            document.body.classList.toggle("theme-dark")
                        }}></span></label>
                    </div>
                </div>
                
            </div>
        </div>
        <a href="/main/logout"class="p-2 text-center ms-3 menu-icon "><i class="btn-round-md bg-red-gradiant text-white feather feather-log-out font-md me-3"></i></a>

        <a href="/main/profile" class="p-0 ms-3 menu-icon"><img src="/images/profile-4.png" alt="user" class="w40 mt--1"></a>
        
    </div>
    
