<script>
	import { ContactsStore } from '../stores';
	export let data;
	import { PUBLIC_BACKEND_URL_FRONTEND } from '$env/static/public';
	let success = true;
	let erroralert = true;

	import axios from "axios";

	async function RequestFollow(id_user) {
		try {
			let header = {
				cookie: document.cookie,
			};
			const config = {
				method: "get",
				withCredentials: true,
				header,
				mode: "no-cors",
				params: { user_id: id_user },
			};
			let response = await axios(PUBLIC_BACKEND_URL_FRONTEND+"/follow", config);
			if (response?.data?.success) {
				success = false;
				setTimeout(() => {
					success = true;
				}, 2000);

				data.listusers = data.listusers.filter((user) => user.ID != id_user);
			} else {
				erroralert = false;
				setTimeout(() => {
					erroralert = true;
				}, 2000);
			}
		} catch (err) {
			erroralert = true;
			console.log("follow aasmmasm",err);
			setTimeout(() => {
				erroralert = false;
			}, 2000);
		}
	}
</script>

{#key data}
	<div class="main-content right-chat-active">
		<div class="container">
			<!-- Alert Success -->
			<div
				id="success-alert"
				class="alert alert-success alert-dismissible {success
					? 'd-none'
					: ''} fade show"
				role="alert"
			>
				friend request success!
			</div>
			<!-- Alert Error -->
			<div
				id="error-alert"
				class="alert alert-danger alert-dismissible {erroralert
					? 'd-none'
					: ''} fade show"
				role="alert"
			>
				Error for request friend!
			</div>
		</div>
		<div class="middle-sidebar-bottom">
			<div class="middle-sidebar-left pe-0">
				<div class="row">
					<div class="col-xl-12">
						<div class="card shadow-xss w-100 d-block d-flex border-0 p-4 mb-3">
							<div class="card-body d-flex align-items-center p-0">
								<h2 class="fw-700 mb-0 mt-0 font-md text-grey-900">Friends</h2>
								<div class="search-form-2 ms-auto">
									<i class="ti-search font-xss"></i>
									<input
										type="text"
										class="form-control text-grey-500 mb-0 bg-greylight theme-dark-bg border-0"
										placeholder="Search here."
									/>
								</div>
								<a
									href="#"
									class="btn-round-md ms-2 bg-greylight theme-dark-bg rounded-3"
									><i class="feather-filter font-xss text-grey-500"></i></a
								>
							</div>
						</div>
						<div class="row ps-2 pe-2">
							{#each data.listusers as user}
								<div class="col-md-3 col-sm-4 pe-2 ps-2">
									<div
										class="card d-block border-0 shadow-xss rounded-3 overflow-hidden mb-3"
									>
										<div
											class="card-body d-block w-100 ps-3 pe-3 pb-4 text-center"
										>
											<figure
												class="avatar ms-auto me-auto mb-0 position-relative w-50 z-index-1"
											>
												{#if user.Avatar}
													<img
														src={user.Avatar}
														alt={user.FirstName + " " + user.LastName}
														class="bg-white rounded-circle" style="width: 70px; height: 70px;"
													/>
												{:else}
													<img
														src="//ui-avatars.com/api/?name={user.FirstName +
															' ' +
															user.LastName}e&size=70&rounded=true&color=fff&background=random"
														alt={user.FirstName + " " + user.LastName}
														class="bg-white rounded-circle shadow-xss"
													/>
												{/if}
											</figure>
											{#if user.IsPublic === 1}
												<a href="/main/profile/{user.ID}">
													<h4 class="fw-700 font-xsss mt-3 mb-1">
														{user.FirstName + " " + user.LastName}
													</h4>
												</a>
											{:else}
												<h4 class="fw-700 font-xsss mt-3 mb-1">
													{user.FirstName + " " + user.LastName}
												</h4>
											{/if}
											<p class="fw-500 font-xsssss text-grey-500 mt-0 mb-3">
												{user.Email}
											</p>
											<a
												href="#"
												on:click={RequestFollow(user.ID)}
												class="mt-0 btn pt-2 pb-2 ps-3 pe-3 lh-24 ms-1 ls-3 d-inline-block rounded-xl bg-success font-xsssss fw-700 ls-lg text-white"
												>ADD FRIEND</a
											>
										</div>
									</div>
								</div>
							{/each}
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
{/key}
