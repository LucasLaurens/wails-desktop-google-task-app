<script lang="ts" setup>
import { ref } from 'vue';
import * as runtime from '../../wailsjs/runtime'
import { WailsEmitTaskWithDescription } from '../../wailsjs/go/main/App';

const task = ref<string>("")

function createTask() {
  console.log("create new task", task.value)
  runtime.EventsEmit('createTask', task.value)
  task.value = ""
}
</script>

<template>
    <div>
        <h1 style="color: white;">Create a new task</h1>
        <input
          v-model="task"
          placeholder="Enter a new task..."
          type="text"
          name="task"
          id="task"
          @keyup.enter="createTask"
        />
    </div>
</template>