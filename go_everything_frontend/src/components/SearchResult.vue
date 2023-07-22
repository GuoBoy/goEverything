<template>
    <table>
        <thead>
            <tr>
                <th>ID</th>
                <th>Name</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="item in items" :key="item.id">
                <td>{{ item.id }}</td>
                <td>{{ item.name }}</td>
            </tr>
        </tbody>
    </table>
</template>

<script lang="ts" setup>
interface Item {
    id: string
    name: string
    type: number
}

const items: Item[] = [{ id: "hello", name: "我是文件", type: 0 }, { id: "word", name: "我是文件夹", type: 1 }]
</script>

<style lang="scss" scoped>
table,
td {
    border-left: 1px solid #333;
}

thead,
tfoot {
    background-color: #333;
    color: #fff;
}
</style>