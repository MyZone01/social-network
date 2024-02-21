<template>
  <NuxtLayout>
    <main id="site__main" class="2xl:ml-[--w-side]  xl:ml-[--w-side-sm] p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top]">
      <div class="max-w-[1065] mx-auto bg-slate-500">
        <div class="bg-white shadow lg:rounded-b-2xl lg:-mt-10 dark:bg-dark2">
          <div class="relative overflow-hidden w-full lg:h-72 h-32">
            <img src="assets/images/post/img-2.jpg" alt="" class="h-full w-full object-cover inset-0">
            <div class="w-full bottom-0 absolute left-0 bg-gradient-to-t from-black/60 pt-10 z-10"></div>
            <div class="absolute bottom-0 right-0 m-a z-20">
              <div class="flex items-center gap-3 px-2">
                <UButton class="button bg-white/20 text-white flex items-center gap-2 back-drop-blur-small">Crop</UButton>
                <UButton class="button bg-white/10 text-white flex items-center gap-2 back-drop-blur-small">Edit</UButton>

              </div>
            </div>
          </div>

          <div class="lg:px-10 md:p-5 p-3">
            <div class="flex flex-col justify-center">
              <div class="flex lg:items-center justify-between max-md:flex-col">
                <div class="flex-1">
                  <h3 class="md:text-2xl "></h3>
                </div>
              </div>
            </div>
          </div>
        </div>


      </div>
    </main>
  </NuxtLayout>
</template>

<script lang="ts" setup>

</script>

<style></style>