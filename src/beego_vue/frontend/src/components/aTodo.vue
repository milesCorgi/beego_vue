<template>
    <div class="home">
        <div style="margin-top: 10px;margin-bottom: 10px">
            <el-row>
                <el-input v-model="input" placeholder="新Todo" style="display:inline-table; width: 30%"></el-input>
                <el-button type="primary" @click="addTodo">新增1个ToDo</el-button>
            </el-row>
            <el-row>
                <div class="block">
                    <el-date-picker
                            v-model="SearchDate"
                            align="right"
                            type="datetimerange"
                            value-format="yyyy-MM-dd HH:mm:ss"
                            range-separator="至"
                            start-placeholder="开始日期"
                            end-placeholder="结束日期"
                            :picker-options="pickerOptions">
                    </el-date-picker>
                    <el-button type="primary" @click="searchTodo">搜索</el-button>
                </div>
            </el-row>
        </div>

        <div>
            <el-row>
                <el-table :data="todoList" style="width: 100%" border>
                  <el-table-column prop="book_name" label="编号" min-width="100">
                    <template scope="scope"> {{ scope.row.Id }}</template>
                  </el-table-column>
                    <el-table-column prop="book_name" label="Todo" min-width="100">
                        <template scope="scope"> {{ scope.row.Todo_body }}</template>
                    </el-table-column>
                    <el-table-column prop="add_time" label="添加时间" min-width="100">
                        <template scope="scope"> {{ scope.row.Add_time }}</template>
                    </el-table-column>
                    <el-table-column prop="id" label="更新时间" min-width="100">
                        <template scope="scope"> {{ scope.row.Update_time }}</template>
                    </el-table-column>
                    <el-table-column prop="id" label="状态" min-width="100">
                        <template scope="scope"> {{ scope.row.Status }}</template>
                    </el-table-column>
                    <el-table-column label="操作" min-width="100">
                        <template scope="scope">
                            <i class="el-icon-edit" @click="openEditDialog(scope.row.Id)"
                               style="cursor:pointer;margin-right: 20px">编辑</i>
                            <i class="el-icon-delete" @click="delete_todo(scope.row.Id)" style="cursor:pointer"
                               scope="scope">删除</i>
                        </template>

                    </el-table-column>
                </el-table>
            </el-row>
        </div>

        <el-dialog
                :visible.sync="openEditDialogFlag"
                @close="openEditDialogFlag = false">
            <el-input v-model="toBeEditBody" type="input" style="display:inline-table; width: 50%"></el-input>
            <el-button type="primary" @click="editTodo">修改</el-button>
        </el-dialog>
    </div>
</template>

<script>
import Qs from 'qs'

export default {
  name: 'home',
  data () {
    return {
      input: '',
      todoList: [],
      openEditDialogFlag: false,
      toBeEditId: 0,
      toBeEditBody: '',
      SearchDate: []
    }
  },
  mounted: function () {
    this.showtodos()
  },
  methods: {
    addTodo () {
      this.$http.post('/api/add_todos', {'Todo_body': this.input}, {emulateJSON: true})
        .then((response) => {
          console.log(response)
          let res = response.data
          if (res.error_num === 0) {
            this.showtodos()
            this.input = ''
          } else {
            alert('新增Todo失败，请重试')
            console.log(res['msg'])
          }
        })
    },
    searchTodo () {
      let posData = Qs.stringify({'search_range': this.SearchDate})
      this.$http.post('/api/search_todos', posData)
        .then((response) => {
          console.log(response)
          let res = response.data
          console.log(res)
          if (res.error_num === 0) {
            this.todoList = res['list']
            this.input = ''
          } else {
            alert('查询Todo失败，请重试')
            // 查询失败的时候，默认先显示全部的
            this.showtodos()
            console.log(res['msg'])
          }
        })
    },
    showtodos () {
      this.$http.get('/api/show_todos')
        .then((response) => {
          let res = response.data
          console.log(res)
          if (res.error_num === 0) {
            this.todoList = res['list']
            console.log(this.todoList)
          } else {
            this.$message.error('查询Todo失败')
            console.log(res['msg'])
          }
        })
    },
    delete_todo (id) {
      this.$http.post('/api/delete_todos', {Todo_id: id}, {emulateJSON: true})
        .then((response) => {
          console.log(response)
          let res = response.data
          if (res.error_num === 0) {
            this.showtodos()
          } else {
            alert('删除Todo失败，请重试')
            console.log(res['msg'])
          }
        })
    },
    editTodo () {
      this.$http.post('/api/edit_todos', {Todo_id: this.toBeEditId, Todo_body: this.toBeEditBody}, {emulateJSON: true})
        .then((response) => {
          console.log(response)
          let res = response.data
          if (res.error_num === 0) {
            this.openEditDialogFlag = false
            this.showtodos()
          } else {
            alert('新增Todo失败，请重试')
            console.log(res['msg'])
          }
        })
    },
    openEditDialog (id) {
      this.openEditDialogFlag = true
      this.toBeEditId = id
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    h1, h2 {
        font-weight: normal;
    }

    ul {
        list-style-type: none;
        padding: 0;
    }

    li {
        display: inline-block;
        margin: 0 10px;
    }

    a {
        color: #42b983;
    }
</style>
