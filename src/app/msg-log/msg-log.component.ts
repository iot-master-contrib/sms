 
 import { RequestService } from '../request.service';
  import { Component, ViewChild } from '@angular/core';
  import { Router } from '@angular/router';
  import { NzMessageService } from 'ng-zorro-antd/message';
  import { NzTableQueryParams } from 'ng-zorro-antd/table';
  import { ParseTableQuery } from '../base/table';
  @Component({
    selector: 'app-msg-log',
    templateUrl: './msg-log.component.html',
    styleUrls: ['./msg-log.component.scss']
  })
  export class MsgLogComponent {
    constructor(
      private router: Router,
      private rs: RequestService,
      private msg: NzMessageService
    ) {
      this.load();
    }
  
    addVisible = false;
    loading = true;
    datum: any[] = [];
    total = 1;
    pageSize = 20;
    pageIndex = 1;
    query: any = {};
    load() {
      this.loading = true;
      this.rs
        .post('sms/search', this.query)
        .subscribe((res) => {
          this.datum = res.data;
          this.total = res.total;
        })
        .add(() => {
          this.loading = false;
        });
    }
    delete(index: number, id: number) {
      this.rs.get(`sms/${id}/delete`).subscribe((res) => {
        this.msg.success('删除成功');
        this.load();
      });
    }
     
     
    onQuery($event: NzTableQueryParams) {
      ParseTableQuery($event, this.query);
      this.load();
    }
    pageIndexChange(pageIndex: number) {
      this.query.skip = pageIndex - 1;
    }
    pageSizeChange(pageSize: number) {
      this.query.limit = pageSize;
    }
    search(text: any) {
      if (text)
        this.query.filter = {
          id: text,
        };
      else this.query = {};
      this.load();
    }
    cancel() {
      this.msg.info('取消删除');
    }
  
    open(id: string) {
      //this.router.navigateByUrl("/admin/device/" + id)
    }
  
  }
  
