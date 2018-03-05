// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: engine.proto
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Pulumirpc {

  /// <summary>Holder for reflection information generated from engine.proto</summary>
  public static partial class EngineReflection {

    #region Descriptor
    /// <summary>File descriptor for engine.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static EngineReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CgxlbmdpbmUucHJvdG8SCXB1bHVtaXJwYxobZ29vZ2xlL3Byb3RvYnVmL2Vt",
            "cHR5LnByb3RvIkcKCkxvZ1JlcXVlc3QSKAoIc2V2ZXJpdHkYASABKA4yFi5w",
            "dWx1bWlycGMuTG9nU2V2ZXJpdHkSDwoHbWVzc2FnZRgCIAEoCSo6CgtMb2dT",
            "ZXZlcml0eRIJCgVERUJVRxAAEggKBElORk8QARILCgdXQVJOSU5HEAISCQoF",
            "RVJST1IQAzJACgZFbmdpbmUSNgoDTG9nEhUucHVsdW1pcnBjLkxvZ1JlcXVl",
            "c3QaFi5nb29nbGUucHJvdG9idWYuRW1wdHkiAGIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Google.Protobuf.WellKnownTypes.EmptyReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::Pulumirpc.LogSeverity), }, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Pulumirpc.LogRequest), global::Pulumirpc.LogRequest.Parser, new[]{ "Severity", "Message" }, null, null, null)
          }));
    }
    #endregion

  }
  #region Enums
  /// <summary>
  /// LogSeverity is the severity level of a log message.  Errors are fatal; all others are informational.
  /// </summary>
  public enum LogSeverity {
    /// <summary>
    /// a debug-level message not displayed to end-users (the default).
    /// </summary>
    [pbr::OriginalName("DEBUG")] Debug = 0,
    /// <summary>
    /// an informational message printed to output during resource operations.
    /// </summary>
    [pbr::OriginalName("INFO")] Info = 1,
    /// <summary>
    /// a warning to indicate that something went wrong.
    /// </summary>
    [pbr::OriginalName("WARNING")] Warning = 2,
    /// <summary>
    /// a fatal error indicating that the tool should stop processing subsequent resource operations.
    /// </summary>
    [pbr::OriginalName("ERROR")] Error = 3,
  }

  #endregion

  #region Messages
  public sealed partial class LogRequest : pb::IMessage<LogRequest> {
    private static readonly pb::MessageParser<LogRequest> _parser = new pb::MessageParser<LogRequest>(() => new LogRequest());
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<LogRequest> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Pulumirpc.EngineReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public LogRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public LogRequest(LogRequest other) : this() {
      severity_ = other.severity_;
      message_ = other.message_;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public LogRequest Clone() {
      return new LogRequest(this);
    }

    /// <summary>Field number for the "severity" field.</summary>
    public const int SeverityFieldNumber = 1;
    private global::Pulumirpc.LogSeverity severity_ = 0;
    /// <summary>
    /// the logging level of this message.
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Pulumirpc.LogSeverity Severity {
      get { return severity_; }
      set {
        severity_ = value;
      }
    }

    /// <summary>Field number for the "message" field.</summary>
    public const int MessageFieldNumber = 2;
    private string message_ = "";
    /// <summary>
    /// the contents of the logged message.
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Message {
      get { return message_; }
      set {
        message_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as LogRequest);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(LogRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Severity != other.Severity) return false;
      if (Message != other.Message) return false;
      return true;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Severity != 0) hash ^= Severity.GetHashCode();
      if (Message.Length != 0) hash ^= Message.GetHashCode();
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Severity != 0) {
        output.WriteRawTag(8);
        output.WriteEnum((int) Severity);
      }
      if (Message.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(Message);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Severity != 0) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) Severity);
      }
      if (Message.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Message);
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(LogRequest other) {
      if (other == null) {
        return;
      }
      if (other.Severity != 0) {
        Severity = other.Severity;
      }
      if (other.Message.Length != 0) {
        Message = other.Message;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            input.SkipLastField();
            break;
          case 8: {
            severity_ = (global::Pulumirpc.LogSeverity) input.ReadEnum();
            break;
          }
          case 18: {
            Message = input.ReadString();
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code
