<template lang="">
    <div id="create-event-overlay" class="hidden lg:p-20 uk- open" uk-modal="">
      <div
        class="uk-modal-dialog tt relative overflow-hidden mx-auto bg-white shadow-xl rounded-lg md:w-[520px] w-full dark:bg-dark2"
      >
        <div class="text-center py-4 border-b mb-0 dark:border-slate-700">
          <h2 class="text-sm font-medium text-black">
            Create Event
          </h2>
  
          <!-- close button -->
          <button type="button" class="button-icon absolute top-0 right-0 m-2.5 uk-modal-close">
            <svg
              xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
              stroke="currentColor" class="w-6 h-6"
            >
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
  
        <form ref="create_event_form" action="">
          <div class="space-y-5 mt-3 p-2">
            <input
              id="" label="Title"
              class="w-full !text-black placeholder:!text-black !bg-white !border-transparent focus:!border-transparent focus:!ring-transparent !font-normal !text-xl   dark:!text-white-100 dark:placeholder:!text-red-500 dark:!bg-slate-800" name="title" type="text" placeholder="title"
            />
            <textarea id="" class="resize-none w-full" cols="30" rows="5" name="description" placeholder="Description" />
            <input name="date_time" placeholder="01/01/2000" label="Date" type="datetime-local" required />
          </div>
  
  
          <div class="p-5 flex justify-between items-center">
            <div class="flex items-center gap-2">
              <button type="submit" class="button bg-blue-500 text-white py-2 px-12 text-[14px]">
                Create
              </button>
            </div>
          </div>
        </form>
      </div>
    </div>
  </template>

<script setup>
const create_event_form = ref(null)
const { createEvent } = useEvents()
const props = defineProps({
  groupId: String,
});
onMounted(() => {
  create_event_form.value?.addEventListener("submit", submitData);
});
async function submitData(e) {
  e.preventDefault();
  const formData = new FormData(e.target)
  const data = Object.fromEntries(formData.entries())
  data.date_time = new Date(data.date_time)

  const response = await createEvent(data, props.groupId)

  

}
</script>

<style lang="">
</style>