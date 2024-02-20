<script>
export default {
    mounted() {

        $(document).ready(function () {
            $('.js-example-basic-multiple').select2({
                placeholder: 'Select friend ..',
            })
        })
    }
}

</script>

<template>
  <select id="selecUser" class="js-example-basic-multiple  " style="width: 40%" name="states[]" multiple="multiple">
    <option value="AL">
      Alabama
    </option>

    <option value="WY">
      Wyoming
    </option>
    <option value="WY">
      Wyoming
    </option>
    <option value="WY">
      Wyoming
    </option>
    <option value="WY">
      Wyoming
    </option>
    <option value="WY">
      Wyoming
    </option>
    <option value="WY">
      Wyoming
    </option>
    <option value="WY">
      Wyoming
    </option>
    <option value="WY">
      Wyoming
    </option>
    <option value="WY">
      Wyoming
    </option>
    <option value="WY">
      Wyoming
    </option>
    <option value="WY">
      Wyoming
    </option>
  </select>
</template>
<style>
.select2-dropdown {
    background-color: rgb(30 41 59 / var(--tw-bg-opacity));
    border: none !important;
}

.select2-container--default .select2-selection--multiple {
    background-color: rgb(51 65 85 / 1);
    border: none !important;
    max-height: 8vh;
    overflow-y: auto;
    overflow-x: hidden;

}
.select2-container--default .select2-selection--multiple .select2-selection__choice {
    background-color: rgb(30 41 59 / var(--tw-bg-opacity));
    border: none !important;
}

.select2-container--default .select2-results__option--selected {
    background-color: rgb(51 65 85 / 1);
}
</style>

