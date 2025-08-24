import { Component,signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ChatService } from '../../services/chat';

@Component({
  selector: 'app-chat',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './chat.html',
  styleUrls: ['./chat.css']
})
export class ChatComponent {
  username : string = '';
  message  : string = '';
  messages =  signal<string[]>([]);
  connected = false;
  onlineCount =  signal<number>(0);

  constructor(private chatService: ChatService) {
    this.chatService.messages$.subscribe(msg => {
      const temp = [...this.messages(),msg]
      this.messages.set(temp);
    });
    this.refreshOnlineCount();
    setInterval(() => this.refreshOnlineCount(), 3000);
  }



private async refreshOnlineCount(): Promise<void> {
  try {
    const n = await this.chatService.getOnlineCount();
    this.onlineCount.set(n);
  } catch (e) {
    console.error('can t get count', e);
  }
}

  connect() {
  if (this.username.trim()) {
    this.chatService.connect(this.username);
     this.connected = true;
  }
}

  send() {
    if (this.message.trim()) {
      this.chatService.sendMessage(this.message);
      this.message = '';
    }
  }
}
