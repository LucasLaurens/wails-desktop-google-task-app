<script lang="ts" setup>
import { ref } from 'vue';
import * as runtime from '../../wailsjs/runtime'

const task = ref<string>("")

function createTask() {
  console.log("create new task", task.value)
  runtime.EventsEmit('createTask', task.value)
  task.value = ""
}
</script>

<template>
  <div class="create-task">
    <div class="create-task-container">
      <input
        v-model="task"
        placeholder="Enter a new task..."
        type="text"
        name="task"
        id="task"
        class="create-task-input"
        @keyup.enter="createTask"
      />
    </div>
  </div>
</template>

<style scoped>
.create-task {
  width: 100%;
  height: 100%;
  overflow: hidden;
}

.create-task-container {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2px;
}

.create-task-input {
  width: calc(100% - 25px);
  height: calc(100% - 25px);
  background-color: transparent;
  padding: 2px;
  border: none;
  outline: none;
  color: white;
  font-size: 18px;
}
</style>