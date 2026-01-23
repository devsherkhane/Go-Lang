<script setup>
defineProps({
  courses: Array,
});

const emit = defineEmits(["edit-course", "delete-course"]);
</script>

<template>
  <h2>Course List</h2>

  <table border="1">
    <tr>
      <th>Name</th>
      <th>Price</th>
      <th>Author</th>
      <th>Actions</th>
    </tr>

    <tr v-for="course in courses" :key="course.courseid">
      <td>{{ course.coursename }}</td>
      <td>{{ course.courseprice }}</td>
      <td>{{ course.author }}</td>
      <td>
        <button @click="emit('edit-course', course)">Edit</button>
        <button @click="emit('delete-course', course.courseid)">Delete</button>
      </td>
    </tr>
  </table>
</template>
<style scoped>
h2 {
  color: #1e293b;
  margin-bottom: 1.5rem;
  font-weight: 700;
}

.table-container {
  overflow-x: auto;
}

table {
  width: 100%;
  border-collapse: separate;
  border-spacing: 0 8px; /* Creates the "pill" row look */
}

th {
  text-align: left;
  padding: 12px 20px;
  color: #64748b;
  font-size: 0.85rem;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

tr {
  background: white;
  transition: transform 0.2s ease;
}

td {
  padding: 16px 20px;
  border-top: 1px solid #f1f5f9;
  border-bottom: 1px solid #f1f5f9;
}

td:first-child {
  border-left: 1px solid #f1f5f9;
  border-top-left-radius: 12px;
  border-bottom-left-radius: 12px;
  font-weight: 600;
  color: #4338ca;
}

td:last-child {
  border-right: 1px solid #f1f5f9;
  border-top-right-radius: 12px;
  border-bottom-right-radius: 12px;
}

.actions {
  display: flex;
  gap: 10px;
}

.btn-edit {
  background: #f5f3ff;
  color: #7c3aed;
  border: none;
  padding: 8px 16px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
}

.btn-delete {
  background: #fff1f2;
  color: #e11d48;
  border: none;
  padding: 8px 16px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
}

button:hover {
  filter: brightness(0.95);
  transform: scale(1.05);
}
</style>