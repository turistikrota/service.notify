package notify

import (
	"reflect"
	"testing"

	"github.com/mixarchitecture/i18np"
)

func Test_newNotifyErrors(t *testing.T) {
	tests := []struct {
		name string
		want Errors
	}{
		{
			name: "newNotifyErrors",
			want: &notifyErrors{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newNotifyErrors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newNotifyErrors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_notifyErrors_Failed(t *testing.T) {
	type args struct {
		operation string
	}
	tests := []struct {
		name string
		e    *notifyErrors
		args args
		want *i18np.Error
	}{
		{
			name: "Failed",
			e:    &notifyErrors{},
			args: args{
				operation: "test",
			},
			want: i18np.NewError(I18nMessages.Failed, i18np.P{"Operation": "test"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Failed(tt.args.operation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("notifyErrors.Failed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_notifyErrors_TypeNotFound(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		e    *notifyErrors
		args args
		want *i18np.Error
	}{
		{
			name: "TypeNotFound",
			e:    &notifyErrors{},
			args: args{
				t: "test",
			},
			want: i18np.NewError(I18nMessages.TypeNotFound, i18np.P{"Type": "test"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.TypeNotFound(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("notifyErrors.TypeNotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_notifyErrors_ValidationFailed(t *testing.T) {
	type args struct {
		operation string
	}
	tests := []struct {
		name string
		e    *notifyErrors
		args args
		want *i18np.Error
	}{
		{
			name: "ValidationFailed",
			e:    &notifyErrors{},
			args: args{
				operation: "test",
			},
			want: i18np.NewError(I18nMessages.ValidationFailed, i18np.P{"Operation": "test"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ValidationFailed(tt.args.operation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("notifyErrors.ValidationFailed() = %v, want %v", got, tt.want)
			}
		})
	}
}
