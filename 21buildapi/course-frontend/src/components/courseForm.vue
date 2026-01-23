<script setup>
import { ref, watch } from "vue";

const props = defineProps({
  course: Object,
  mode: String,
});

const emit = defineEmits(["submit-course"]);

const form = ref({
  coursename: "",
  courseprice: "",
  author: "",
});

// When editing, fill form
watch(
  () => props.course,
  (newCourse) => {
    if (newCourse) {
      form.value = { ...newCourse };
    } else {
      form.value = { coursename: "", courseprice: "", author: "" };
    }
  },
  { immediate: true }
);

function submitForm() {
  emit("submit-course", {
    coursename: form.value.coursename,
    courseprice: form.value.courseprice,
    author: form.value.author,
  });
}
</script>

<template>
  <h2>{{ mode === "create" ? "Add Course" : "Edit Course" }}</h2>

  <input v-model="form.coursename" placeholder="Course Name" color="black" />
  <input v-model="form.courseprice" placeholder="Price" color="black"/>
  <input v-model="form.author" placeholder="Author" color="black"/>

  <button @click="submitForm">
    {{ mode === "create" ? "Create" : "Update" }}
  </button>
</template>
  <style scoped>
h2 {
  color: #4338ca;
  margin-bottom: 1.5rem;
  font-weight: 700;
}

.form-layout {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1.5rem;
  align-items: flex-end;
}

.input-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

input {
  padding: 12px 16px;
  border: 2px solid #e2e8f0;
  border-radius: 12px;
  font-size: 1rem;
  transition: all 0.3s ease;
  background: rgb(0, 0, 0);
}

input:focus {
  outline: none;
  border-color: #6366f1;
  box-shadow: 0 0 0 4px rgba(99, 102, 241, 0.15);
}

.btn-primary {
  background: #6366f1;
  color: white;
  padding: 12px 24px;
  border: none;
  border-radius: 12px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 6px rgba(99, 102, 241, 0.25);
}

.btn-primary:hover {
  background: #4f46e5;
  transform: translateY(-2px);
  box-shadow: 0 6px 12px rgba(99, 102, 241, 0.3);
}
</style>