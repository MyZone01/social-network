<template>
    <div class="hidden lg:p-20" id="create-story" uk-modal="">

        <div
            class="uk-modal-dialog tt relative overflow-hidden mx-auto bg-white p-7 shadow-xl rounded-lg md:w-[520px] w-full dark:bg-dark2">

            <div class="text-center py-3 border-b -m-7 mb-0 dark:border-slate-700">
                <h2 class="text-sm font-medium"> Create Status </h2>

                <!-- close button -->
                <button type="button" class="button__ico absolute top-0 right-0 m-2.5 uk-modal-close">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                        stroke="currentColor" class="w-6 h-6">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                </button>

            </div>

            <div class="space-y-5 mt-7">

                <div>
                    <label for="" class="text-base">What do you have in mind? </label>
                    <input type="text" class="w-full mt-3">
                </div>

                <div>
                    <div
                        class="w-full h-72 relative border1 rounded-lg overflow-hidden bg-[url('../images/ad_pattern.png')] bg-repeat">

                        <label for="createStatusUrl"
                            class="flex flex-col justify-center items-center absolute -translate-x-1/2 left-1/2 bottom-0 z-10 w-full pb-6 pt-10 cursor-pointer bg-gradient-to-t from-gray-700/60">
                            <input id="createStatusUrl" type="file" class="hidden" />
                            <ion-icon name="image" class="text-3xl text-teal-600"></ion-icon>
                            <span class="text-white mt-2">Browse to Upload image </span>
                        </label>

                        <img id="createStatusImage" src="#" alt="Uploaded Image" accept="image/png, image/jpeg"
                            style="display:none;" class="w-full h-full absolute object-cover">

                    </div>

                </div>

                <div class="flex justify-between items-center">

                    <div class="flex items-start gap-2">
                        <ion-icon name="time-outline"
                            class="text-3xl text-sky-600  rounded-full bg-blue-50 dark:bg-transparent"></ion-icon>
                        <p class="text-sm text-gray-500 font-medium"> Your Status will be available <br> for <span
                                class="text-gray-800"> 24 Hours</span> </p>
                    </div>

                    <button type="button" class="button bg-blue-500 text-white px-8"> Create</button>

                </div>

            </div>

        </div>

    </div>
</template>