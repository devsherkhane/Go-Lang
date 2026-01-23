<script setup>
import { ref, onMounted } from "vue";
import CourseForm from "./components/CourseForm.vue";
import CourseList from "./components/CourseList.vue";
import {
  getCourses,
  createCourse,
  updateCourse,
  deleteCourse,
} from "./services/courseService";

const courses = ref([]);
const selectedCourse = ref(null);
const mode = ref("create");

// Load data
onMounted(async () => {
  courses.value = await getCourses();
});

// Create or Update
async function handleSubmit(course) {
  if (mode.value === "create") {
    const res = await createCourse(course);
    courses.value.push({ ...course, courseid: res.id });
  } else {
    await updateCourse(selectedCourse.value.courseid, course);
    const index = courses.value.findIndex(
      c => c.courseid === selectedCourse.value.courseid
    );
    courses.value[index] = { ...course, courseid: selectedCourse.value.courseid };
    mode.value = "create";
    selectedCourse.value = null;
  }
}

// Edit
function handleEdit(course) {
  selectedCourse.value = course;
  mode.value = "edit";
}

// Delete
async function handleDelete(id) {
  await deleteCourse(id);
  courses.value = courses.value.filter(c => c.courseid !== id);
}
</script>

<template>
  <h1>Course Management</h1>

  <CourseForm
    :course="selectedCourse"
    :mode="mode"
    @submit-course="handleSubmit"
  />

  <CourseList
    :courses="courses"
    @edit-course="handleEdit"
    @delete-course="handleDelete"
  />
</template>
<style>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;600;800&display=swap');

body {
  margin: 0;
  font-family: 'Plus Jakarta Sans', sans-serif;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  min-height: 100vh;
  padding: 2rem;
  color: #1e293b;
}

.container {
  max-width: 1000px;
  margin: 0 auto;
}

h1 {
  font-size: 2.5rem;
  font-weight: 800;
  text-align: center;
  color: #1e1b4b;
  margin-bottom: 2rem;
  text-shadow: 0 2px 4px rgba(0,0,0,0.05);
}

.card {
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  border-radius: 20px;
  padding: 2rem;
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.1), 0 8px 10px -6px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.5);
  margin-bottom: 2rem;
}
</style>