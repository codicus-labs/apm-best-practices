---
description: Vue single-file components use script setup with TypeScript
load: on-demand
paths: ["**/*.vue"]
---

**Must.** Every component uses `<script setup lang="ts">`, with typed props through `defineProps<{ order: Order }>()`, events through `defineEmits<{ saved: [id: string] }>()` and two-way bindings through `defineModel()`. No Options API, no `defineComponent`.
