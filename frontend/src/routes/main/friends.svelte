<script>
	export let friends;
</script>

<div class="card w-100 shadow-xss rounded-xxl border-0 mb-3">
	<div class="card-body d-flex align-items-center p-4">
		<h4 class="fw-700 mb-0 font-xssss text-grey-900">Friend Request</h4>
		<a href="/main/follow" class="fw-600 ms-auto font-xssss text-primary"
			>See all</a
		>
	</div>
	{#if friends}
		{#each friends as friend}
			<div
				style="flex-wrap: wrap;"
				class="card-body d-flex align-items-center pt-0 ps-4 pe-4 pb-4"
			>
				<div class="card-body d-flex pt-4 ps-4 pe-4 pb-0 border-top-xs bor-0">
					<figure class="avatar me-3 w60">
						{#if friend.Avatar}
							<img
								src={friend.Avatar}
								alt={friend.FirstName + " " + friend.LastName}
								class="float-right p-0 bg-white rounded-circle w-100 shadow-xss"
							/>
						{:else}
							<img
								src="//ui-avatars.com/api/?name={friend.FirstName +
									' ' +
									friend.LastName}&size=50&rounded=true&color=fff&background=random"
								alt={friend.FirstName + " " + friend.LastName}
							/>
						{/if}
					</figure>
					<h4
						style="margin-top: 10px;"
						class="fw-700 text-grey-900 font-xssss mt-1"
					>
						{friend.FirstName + " " + friend.LastName}
					</h4>
				</div>
			</div>
		{/each}
	{/if}
</div>
