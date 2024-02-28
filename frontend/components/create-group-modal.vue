<template lang="">
    <div class="hidden lg:p-20 uk- open" id="create-group-overlay" uk-modal="">

        <div
            class="uk-modal-dialog tt relative overflow-hidden mx-auto bg-white shadow-xl rounded-lg md:w-[520px] w-full dark:bg-dark2">

            <div class="text-center py-4 border-b mb-0 dark:border-slate-700">
                <h2 class="text-sm font-medium text-black"> Create Group </h2>

                <!-- close button -->
                <button type="button" class="button-icon absolute top-0 right-0 m-2.5 uk-modal-close">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                        stroke="currentColor" class="w-6 h-6">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                </button>

            </div>

            <form action="" ref="create_group_form">

                <div class="space-y-5 mt-3 p-2">
                    <input
                        class="w-full !text-black placeholder:!text-black !bg-white !border-transparent focus:!border-transparent focus:!ring-transparent !font-normal !text-xl   dark:!text-white-100 dark:placeholder:!text-red-500 dark:!bg-slate-800"
                        name="Title" id="" type="text" placeholder="name"/>
                        <textarea class="resize-none w-full" id="" cols="30" rows="10" name="Description" placeholder="Description"></textarea>
                </div>
 

                <div class="p-5 flex justify-between items-center">

                    
                    <div class="flex items-center gap-2">
                        <button type="submit" class="button bg-blue-500 text-white py-2 px-12 text-[14px]">
                            Create</button>
                    </div>
                </div>
            </form>

        </div>


    </div>

</template>
<script>
export default {
    mounted() {
        this.$refs.create_group_form.addEventListener('submit', this.submitData)
    },
    methods: {
        async submitData(e) {
            e.preventDefault()
            const store = useGlobalAuthStore()
            const data = new FormData(e.target)
            const response = await $fetch('/api/group/create', {
                method:'post',
                headers: {
                    Authorization: `Bearer ${store.token}`,
                },
                body: JSON.stringify(Object.fromEntries(data.entries()))
            })

        }
    }
}



</script>
<style lang="">

</style>    