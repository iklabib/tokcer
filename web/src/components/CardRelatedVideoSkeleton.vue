<script setup lang="ts">
import Skeleton from 'primevue/skeleton'
</script>

<template>
    <div class="flex flex-col rounded border border-slate-900 dark:border-slate-400 p-2">
        <Skeleton height="10rem" class="mb-2"></Skeleton>
        <Skeleton height="2rem" class="w-full"></Skeleton>
    </div>
</template>