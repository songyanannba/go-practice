<template>
  <div>
    <div class="gva-search-box">
        <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
           {{- range .Fields}}  {{- if .FieldSearchType}} {{- if eq .FieldType "bool" }}
            <el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">
            <el-select v-model="searchInfo.{{.FieldJson}}" clearable placeholder="Please select">
                <el-option
                    key="true"
                    label="Yes"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="No"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>
           {{- else if .DictType}}
           <el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">
            <el-select v-model="searchInfo.{{.FieldJson}}" clearable placeholder="Please select" @clear="()=>{searchInfo.{{.FieldJson}}=undefined}">
              <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            {{- else}}
        <el-form-item label="{{.FieldDesc}}">
        {{- if eq .FieldType "float64" "int"}}
            {{if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
            <el-input v-model.number="searchInfo.start{{.FieldName}}" placeholder="Search criteria（start）" />
            —
            <el-input v-model.number="searchInfo.end{{.FieldName}}" placeholder="Search criteria（end）" />
           {{- else}}
             {{- if .DictType}}
              <el-select v-model="searchInfo.{{.FieldJson}}" placeholder="Please select" style="width:100%" :clearable="true" >
               <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value" />
             </el-select>
                    {{- else}}
             <el-input v-model.number="searchInfo.{{.FieldJson}}" placeholder="Search criteria" />
                    {{- end }}
          {{- end}}
        {{- else if eq .FieldType "time.Time"}}
            {{if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
            <el-date-picker v-model="searchInfo.start{{.FieldName}}" type="datetime" placeholder="Search criteria（start）"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.end{{.FieldName}}" type="datetime" placeholder="Search criteria（end）"></el-date-picker>
           {{- else}}
           <el-date-picker v-model="searchInfo.{{.FieldJson}}" type="datetime" placeholder="Search criteria"></el-date-picker>
          {{- end}}
        {{- else if eq .FieldType "tinyint"}}
        <el-select v-model="searchInfo.{{.FieldJson}}" :clearable="{{.Clearable}}" filterable placeholder="Please select">
          <el-option v-for="v,k in {{.FieldJson}}Opts" :key="k" :label="v" :value="parseInt(k)"/>
        </el-select>
        {{- else}}
         <el-input v-model="searchInfo.{{.FieldJson}}" placeholder="Search criteria" />
        {{- end}}
        </el-form-item>{{ end }}{{ end }}{{ end }}
        <el-form-item label="Creation time:" prop="betweenTime">
                <el-date-picker
                  v-model="searchInfo.betweenTime"
                  type="datetimerange"
                  :shortcuts="shortcuts"
                  range-separator="To"
                  start-placeholder="Start date"
                  end-placeholder="End date"
                  format="YYYY-MM-DD HH:mm"
                  value-format="YYYYMMDDHHmmss"
                />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">Query</el-button>
          <el-button size="small" icon="refresh" @click="onReset">Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <!-- <el-button size="small" type="primary" icon="download" @click="exportExcel">Export</el-button> -->
            <el-button size="small" type="primary" icon="plus" @click="openDialog">Create</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>Are you sure you want to delete？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="deleteVisible = false">Cancel</el-button>
                <el-button size="small" type="primary" @click="onDelete">Confirm</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">Delete</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        border
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" prop="ID" sortable="custom" width="90" />
        {{- range .Fields}}
        {{- if .DictType}}
        <el-table-column {{- if .Sort}} sortable{{- end}} align="left" label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120">
            <template #default="scope">
            {{"{{"}} filterDict(scope.row.{{.FieldJson}},{{.DictType}}Options) {{"}}"}}
            </template>
        </el-table-column>
        {{- else if eq .FieldType "bool" }}
        <el-table-column {{- if .Sort}} sortable{{- end}} align="left" label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120">
            <template #default="scope">{{"{{"}} formatBoolean(scope.row.{{.FieldJson}}) {{"}}"}}</template>
        </el-table-column>
         {{- else if eq .FieldType "time.Time" }}
         <el-table-column {{- if .Sort}} sortable{{- end}} align="left" label="{{.FieldDesc}}" width="180">
            <template #default="scope">{{"{{"}} formatDate(scope.row.{{.FieldJson}}) {{"}}"}}</template>
         </el-table-column>
         {{- else if eq .FieldType "tinyint" }}
         <el-table-column align="left" label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120">
           <template #default="scope">
             <el-tag size="small" :type="statusTags[scope.row.{{.FieldJson}}]" effect="dark">
               {{"{{"}} {{.FieldJson}}Opts[scope.row.{{.FieldJson}}] {{"}}"}}
             </el-tag>
           </template>
         </el-table-column>
        {{- else }}
        <el-table-column {{- if .Sort}} sortable="custom"{{- end}} align="left" label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120" />
        {{- end }}
        {{- end }}
        <el-table-column align="left" label="Creation time" width="180">
            <template #default="scope">{{ "{{ scope.row.CreatedAt }}" }}</template>
        </el-table-column>
        <el-table-column align="left" label="Update time" width="180">
          <template #default="scope">{{ "{{ scope.row.UpdatedAt }}" }}</template>
        </el-table-column>
        <el-table-column align="left" label="Button group">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="update{{.StructName}}Func(scope.row)">Change</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">Delete</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="Pop-up operation">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="120px">
    {{- range .Fields}}
        <el-form-item label="{{.FieldDesc}}:"  prop="{{.FieldJson}}" >
      {{- if eq .FieldType "bool" }}
          <el-switch v-model="formData.{{.FieldJson}}" active-color="#13ce66" inactive-color="#ff4949" active-text="Yes" inactive-text="No" clearable ></el-switch>
      {{- end }}
      {{- if eq .FieldType "string" }}
          <el-input v-model="formData.{{.FieldJson}}" :clearable="{{.Clearable}}"  placeholder="Please enter" />
      {{- end }}
      {{- if eq .FieldType "int" }}
      {{- if .DictType}}
          <el-select v-model="formData.{{ .FieldJson }}" placeholder="Please select" style="width:100%" :clearable="{{.Clearable}}" >
            <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value" />
          </el-select>
      {{- else }}
          <el-input v-model.number="formData.{{ .FieldJson }}" :clearable="{{.Clearable}}" placeholder="Please enter" />
      {{- end }}
      {{- end }}
      {{- if eq .FieldType "time.Time" }}
          <el-date-picker v-model="formData.{{ .FieldJson }}" type="date" style="width:100%" placeholder="Select date" :clearable="{{.Clearable}}"  />
      {{- end }}
      {{- if eq .FieldType "float64" }}
          <el-input-number v-model="formData.{{ .FieldJson }}"  style="width:100%" :precision="2" :clearable="{{.Clearable}}"  />
      {{- end }}
      {{- if eq .FieldType "enum" }}
            <el-select v-model="formData.{{ .FieldJson }}" placeholder="Please select" style="width:100%" :clearable="{{.Clearable}}" >
               <el-option v-for="item in [{{.DataTypeLong}}]" :key="item" :label="item" :value="item" />
            </el-select>
      {{- end }}
      {{- if eq .FieldType "tinyint" }}
            <el-select v-model="formData.{{ .FieldJson }}" placeholder="Please select" style="width:100%" :clearable="{{.Clearable}}" >
              <el-option v-for="v,k in {{ .FieldJson }}Opts" :key="k" :label="v" :value="parseInt(k)"/>
            </el-select>
      {{- end }}
        </el-form-item>
      {{- end }}
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">Cancel</el-button>
          <el-button size="small" type="primary" @click="enterDialog">Confirm</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: '{{.StructName}}',
  data() {
      return {
        statusTags: { 0: 'warning', 1: 'success', 2: 'info', 3: 'primary', 4: 'danger', 9: 'danger' },
        defaultStatus: { 1: 'Normal', 2: 'Close', 9: 'Abnormal' },
        {{- range .Fields}}
        {{- if eq .FieldType "tinyint"}}
        {{.FieldJson}}Opts: { 1: 'Normal', 2: 'Close'},
        {{- end}}
        {{- end }}
      }
  }
}
</script>

<script setup>
import {
  create{{.StructName}},
  delete{{.StructName}},
  delete{{.StructName}}ByIds,
  update{{.StructName}},
  find{{.StructName}},
  get{{.StructName}}List
} from '@/api/{{.PackageName}}'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { getShortcuts } from '@/utils/date'
import { toSQLLine } from '@/utils/stringFun'
import { exportFile } from '@/utils/excel'

// 自动化生成的字典（可能为空）以及字段
    {{- range $index, $element := .DictTypes}}
const {{ $element }}Options = ref([])
    {{- end }}
const formData = ref({
        {{- range .Fields}}
        {{- if eq .FieldType "bool" }}
        {{.FieldJson}}: false,
        {{- end }}
        {{- if eq .FieldType "string" }}
        {{.FieldJson}}: '',
        {{- end }}
        {{- if eq .FieldType "int" }}
        {{.FieldJson}}: {{- if .DictType }} undefined{{ else }} 0{{- end }},
        {{- end }}
        {{- if eq .FieldType "tinyint" }}
        {{.FieldJson}}: 1,
        {{- end }}
        {{- if eq .FieldType "time.Time" }}
        {{.FieldJson}}: new Date(),
        {{- end }}
        {{- if eq .FieldType "float64" }}
        {{.FieldJson}}: 0,
        {{- end }}
        {{- end }}
        })

// 验证规则
const rule = reactive({
    {{- range .Fields }}
            {{- if eq .Require true }}
               {{.FieldJson }} : [{
                   required: true,
                   message: '{{ .ErrorText }}',
                   trigger: ['input','blur'],
               }],
            {{- end }}
    {{- end }}
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const shortcuts = getShortcuts()

// 排序
const sortChange = ({ prop, order }) => {
  if (prop) {
    searchInfo.value.sort = toSQLLine(prop)
    searchInfo.value.order = order
    getTableData()
  }
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
{{- range .Fields}}{{- if eq .FieldType "bool" }}
  if (searchInfo.value.{{.FieldJson}} === ""){
      searchInfo.value.{{.FieldJson}}=null
  }{{ end }}{{ end }}
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 数据导出
const exportExcel = async() => {
  const map = {
    'ID': '',
    {{- range .Fields}}
    '{{.FieldJson}}': '{{.FieldDesc}}',
    {{- end}}
  }
  const table = await get{{.StructName}}List({ page: 1, pageSize: 100000, ...searchInfo.value })
  if (table.code === 0) {
    const data = formatData(table.data.list)
    const params = {
      sheetData: data,
      headMap: map,
      fileName: 'DataExport'
    }
    exportFile(params)
  } else {
    ElMessage.error(table.msg)
  }
}

// 格式化数据
const formatData = (data) => {
  data.forEach(v => {
    v.CreatedAt = formatDate(v.CreatedAt)
    v.UpdatedAt = formatDate(v.UpdatedAt)
  })
  return data
}

// Query
const getTableData = async() => {
  const table = await get{{.StructName}}List({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = formatData(table.data.list)
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
{{- range $index, $element := .DictTypes }}
    {{ $element }}Options.value = await getDictFunc('{{$element}}')
{{- end }}
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// Delete行
const deleteRow = (row) => {
    ElMessageBox.confirm('Are you sure you want to delete?', 'Prompt', {
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel',
        type: 'warning'
    }).then(() => {
            delete{{.StructName}}Func(row)
        })
    }


// 批量Delete控制标记
const deleteVisible = ref(false)

// 多选Delete
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: 'Please select the data to delete'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await delete{{.StructName}}ByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: 'Delete成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const update{{.StructName}}Func = async(row) => {
    const res = await find{{.StructName}}({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.re{{.Abbreviation}}
        dialogFormVisible.value = true
    }
}


// Delete行
const delete{{.StructName}}Func = async (row) => {
    const res = await delete{{.StructName}}({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: 'Deleted successfully'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
    {{- range .Fields}}
        {{- if eq .FieldType "bool" }}
        {{.FieldJson}}: false,
        {{- end }}
        {{- if eq .FieldType "string" }}
        {{.FieldJson}}: '',
        {{- end }}
        {{- if eq .FieldType "int" }}
        {{.FieldJson}}: {{- if .DictType }} undefined{{ else }} 0{{- end }},
        {{- end }}
        {{- if eq .FieldType "tinyint" }}
        {{.FieldJson}}: 1,
        {{- end }}
        {{- if eq .FieldType "time.Time" }}
        {{.FieldJson}}: new Date(),
        {{- end }}
        {{- if eq .FieldType "float64" }}
        {{.FieldJson}}: 0,
        {{- end }}
        {{- end }}
        }
}
// 弹窗Confirm
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await create{{.StructName}}(formData.value)
                  break
                case 'update':
                  res = await update{{.StructName}}(formData.value)
                  break
                default:
                  res = await create{{.StructName}}(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: 'Saved successfully'
                })
                closeDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>
